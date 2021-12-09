#include <ArduinoOTA.h>

#include <StaticThreadController.h>
#include <Thread.h>
#include <ThreadController.h>

#include <AliyunIoTSDK.h>

#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266httpUpdate.h>


#include <aREST.h>
const char *ssid = "esp8266";
const char *pwd = "esp82666";
static WiFiClient espClient;
const String my_version = "1.06";//版本号，依据这个决定启动时是否需要ota升级
char* my_Version = "1.06";//避免出错
aREST rest = aREST();
#define LISTEN_PORT           80
WiFiServer server(LISTEN_PORT);
Thread my_thread1 = Thread();
Thread my_thread2 = Thread();
//Thread my_thread3 = Thread();
#define PRODUCT_KEY "a1J4LfeKtKw"
#define DEVICE_NAME "Switch1"
#define DEVICE_SECRET "bb90552fb7a551305383a44665dce6f4"
#define REGION_ID "cn-shanghai"
int pre = 0;

int ovalue = 1000;
int cvalue = 2000;


int test = 0;
void go_close(){
             for(int i = 0;i<35;i++){
              digitalWrite(0,1);
              delayMicroseconds(cvalue);
              digitalWrite(0,0);
              delay(20);
           }
}
void go_open(){
        for(int i = 0;i<35;i++){
          digitalWrite(0,1);
          delayMicroseconds(ovalue);
          digitalWrite(0,0);
          delay(20);
        }
}
int opens(String command) {
  go_open();
  Serial.println(command);
  return 1;
}
int closes(String command) {
  go_close();
  Serial.println(command);
  return 0;
}
void change_state(JsonVariant p){
    if(p["IsOpen"] != pre){
      
      pre = p["IsOpen"];
      
      if(p["IsOpen"] == 1){
        Serial.println("open led");
        AliyunIoTSDK::send("IsOpen", 1);
        go_open();
      }else{
           Serial.println("close led");
           AliyunIoTSDK::send("IsOpen", 0);
           go_close();
      }
      
    }
}
void ovalue_change(JsonVariant p){
    //服务端判断值是否合理
    ovalue = p["ovalue"];
    AliyunIoTSDK::send("ovalue",ovalue);
}
void cvalue_change(JsonVariant p){
    //服务端判断值是否合理
    cvalue = p["cvalue"];
    AliyunIoTSDK::send("cvalue",cvalue);
}

void wifiInit(const char *ssid, const char *passphrase)
{
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, passphrase);
    while (WiFi.status() != WL_CONNECTED)
    {
        delay(1000);
        Serial.println("WiFi not Connect");
    }
    Serial.println("Connected to AP");
}
void listen_rest(){
  WiFiClient client = server.available();
  if (!client) {
    return;
  }
  while(!client.available()){
    delay(1);
  }
  rest.handle(client);
}
void listen_ota(){
    ArduinoOTA.handle();
}
void request_ota_check(){
  HTTPClient httpClient;
  httpClient.begin(espClient,"http://ota.dadiqq.cn/switch.txt");
  int httpcode = httpClient.GET();
  Serial.println(httpcode);
  if(httpcode == HTTP_CODE_OK){
    //接收成功,校验版本
    String response = httpClient.getString();
    Serial.println(response);
    if(response != my_version){
        //ota升级
        t_httpUpdate_return ret = ESPhttpUpdate.update(espClient,"ota.dadiqq.cn",80,"/Switch1.bin");
        while(1){
          if(ret != HTTP_UPDATE_OK){
            Serial.println("update fail err -1");
            //出错了,暂时不管，上报一下就行
            AliyunIoTSDK::send("update",false);
            httpClient.end();
            return;
          }else{
            Serial.println("update success");
            AliyunIoTSDK::send("update",true);
            httpClient.end();
            return;
          }  
        }
    }else{
      Serial.println("is new !");
      AliyunIoTSDK::send("update",true);//最新版本无需升级
      httpClient.end();
      return;
    }
  }else{
    Serial.println("update fail err -2");
    AliyunIoTSDK::send("update",false);
    httpClient.end();
    return;
  }
}
void setup(void){
    Serial.begin(115200);
    wifiInit(ssid, pwd);
    rest.set_id("2");
    rest.set_name("esp8266-duoji1");
    rest.function("open",opens);
    rest.function("close",closes);
    rest.variable("test",&test);
    server.begin();
    AliyunIoTSDK::begin(espClient, PRODUCT_KEY, DEVICE_NAME, DEVICE_SECRET, REGION_ID);
    AliyunIoTSDK::bindData("IsOpen", change_state);
    AliyunIoTSDK::bindData("ovalue",ovalue_change);
    AliyunIoTSDK::bindData("cvalue",cvalue_change);
    AliyunIoTSDK::send("FirmwareVersion",my_Version);
    request_ota_check();
    //check完成后上报版本号
    AliyunIoTSDK::send("FirmwareVersion",my_Version);
    pinMode(0,OUTPUT);
    digitalWrite(0,0);
    for(int i = 0;i<35;i++){
      digitalWrite(0,1);
      delayMicroseconds(1500);
      digitalWrite(0,0);
      delay(18);
    }
    ArduinoOTA.begin();
    my_thread1.onRun(listen_rest);
    my_thread1.setInterval(50);
    my_thread2.onRun(listen_ota);
    my_thread2.setInterval(50);
    
}
void loop() {
  // Handle REST calls
  if(my_thread1.shouldRun()){
      my_thread1.run();
  }
  if(my_thread2.shouldRun()){
      my_thread2.run();
  }
  AliyunIoTSDK::loop();
}
