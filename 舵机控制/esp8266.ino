#include <ArduinoOTA.h>

#include <StaticThreadController.h>
#include <Thread.h>
#include <ThreadController.h>

#include <AliyunIoTSDK.h>

#include <ESP8266WiFi.h>
#include <aREST.h>
const char *ssid = "esp8266";
const char *pwd = "esp82666";
static WiFiClient espClient;
aREST rest = aREST();
#define LISTEN_PORT           80
WiFiServer server(LISTEN_PORT);
Thread my_thread1 = Thread();
Thread my_thread2 = Thread();
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
    AliyunIoTSDK::send("IsOpen", pre);
    AliyunIoTSDK::bindData("ovalue",ovalue_change);
    AliyunIoTSDK::bindData("cvalue",cvalue_change);
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
    my_thread1.setInterval(1);
    my_thread2.onRun(listen_ota);
    my_thread2.setInterval(2);
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
