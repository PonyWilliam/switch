<template>
	<view class="container">
		<view class="control" v-for="(item, index) in devices">
			<view class="image">
				<image src="../../static/switch.jpg" mode="aspectFill"></image>
			</view>
			<view class="title">{{item.DeviceName}}</view>
			<view class="sub_title">{{item.Nickname}}</view>
			<view class="state">设备状态:{{item.DeviceStatus == "ONLINE" ? "在线" : "离线"}}</view>
			<view class="switch">
				<view>
				    <switch v-if="open==true" checked @change="switchChange" />
				    <switch v-else-if="open==false" @change="switchChange" />
				</view>
			</view>
			<view class="slider">
				<text class="title">高级设置(oc最高与最低分别是顺逆时针90°)</text>
				开力度:<slider :value="ovalue" block-size="20" block-color="#4CD964" min="0" max="100" show-value="" step="1" backgroundColor="#8080e2" activeColor="orangered" @change="ovalue_change"/>
				关力度:<slider :value="cvalue" block-size="20" block-color="#4CD964" min="0" max="100" show-value="" step="1" backgroundColor="#8080e2" activeColor="orangered" @change="cvalue_change"/>
				<button type="default" class="button" @click="reset_oc">一键复原</button>
			</view>
		</view>
	</view>
</template>

<script>
	const back_url = "https://switchiot.dadiqq.cn"
	export default {
		data(){
			return{
				//返回数据
				devices:[],
				open:false,
				cvalue:0,
				ovalue:0,
			}
		},
		methods: {
			update_device(data){
				let x = 0
				uni.showLoading({
					title:'加载中···'
				})
				this.devices.push(data)
				console.log(this.devices)
				uni.request({
					url:`${back_url}/get/switch`,
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					method:"POST",
					data:{
						"name":this.devices[0].DeviceName,
						"key":this.devices[0].ProductKey
					},
					success:res=>{
						console.log(res.data.data)
						if(res.data.data == "关"){
							this.open = false
						}else{
							this.open = true
						}
						console.log(this.open)
						if(++x>=3){
							uni.hideLoading()
						}
					}
				})
				uni.request({
					url:`${back_url}/get/ocvalue`,
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					method:"POST",
					data:{
						"name":this.devices[0].DeviceName,
						"key":this.devices[0].ProductKey,
						"oc":"o"
					},
					success:res=>{
						this.ovalue = (1500 - res.data.data) / 10;
						console.log(`ovalue:${this.ovalue}`)
						if(++x>=3){
							uni.hideLoading()
						}
					}
				})
				uni.request({
					url:`${back_url}/get/ocvalue`,
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					method:"POST",
					data:{
						"name":this.devices[0].DeviceName,
						"key":this.devices[0].ProductKey,
						"oc":"c"
					},
					success:res=>{
						console.log(res.data.data)
						this.cvalue = (res.data.data - 1500) / 10;
						console.log(`cvalue:${this.cvalue}`)
						if(++x>=3){
							uni.hideLoading()
						}
					}
				})
				
			},
			switchChange(e){
				this.open = e.detail.value
				uni.request({
					url:`${back_url}/set/switch`,
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					method:"POST",
					data:{
						"name":this.devices[0].DeviceName,
						"key":this.devices[0].ProductKey,
						"value":this.open == true ? "1" : "0"
					},
					success:res=>{
						console.log(res)
					}
				})
			},
			ovalue_change(e){
				console.log(`post o:${1500 - (e.detail.value) * 10}`)
				uni.request({
					url:`${back_url}/set/ocvalue`,
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					method:"POST",
					data:{
						"name":this.devices[0].DeviceName,
						"key":this.devices[0].ProductKey,
						"oc":"o",
						"value":1500 - (e.detail.value) * 10
					},
					success:res=>{
						console.log(res)
					}
				})
			},
			cvalue_change(e){
				console.log(`post c:${1500 + (e.detail.value) * 10}`)
				uni.request({
					url:`${back_url}/set/ocvalue`,
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					method:"POST",
					data:{
						"name":this.devices[0].DeviceName,
						"key":this.devices[0].ProductKey,
						"oc":"c",
						"value":1500 + (e.detail.value) * 10
					},
					success:res=>{
				
						console.log(res)
					}
				})
			},
			reset_oc(){
				this.cvalue = 50
				this.ovalue = 50
				let c_config = {
					detail:{
						value:this.cvalue
					}
				}
				let o_config = {
					detail:{
						value:this.ovalue
					}
				}
				this.cvalue_change(c_config)
				this.ovalue_change(o_config)
			}
		},
		onLoad: function(option) {
		  // #ifdef APP-NVUE
		  const eventChannel = this.$scope.eventChannel; // 兼容APP-NVUE
		  // #endif
		  // #ifndef APP-NVUE
		  const eventChannel = this.getOpenerEventChannel();
		  // #endif
		  eventChannel.emit('acceptDataFromOpenedPage', {data: 'data from test page'});
		  eventChannel.emit('someEvent', {data: 'data from test page for someEvent'});
		  // 监听acceptDataFromOpenerPage事件，获取上一页面通过eventChannel传送到当前页面的数据
		  let that = this
		  eventChannel.on('acceptDataFromOpenerPage', function(data) {
		    if(data == (null || undefined)){
				uni.showModal({
					title:'error',
					content:'参数错误',
					showCancel:false
				})
				uni.navigateBack({
					
				})
				return
				
			}
			that.update_device(data.data)
			//访问设备状态
			
		  })
		}
	}
</script>

<style lang="scss">
.container{
	width:100%;
	min-height: 100vh;
	background:#3F536E;
	padding-top:15px;
	.control{
		width:90%;
		box-sizing: border-box;
		display: block;
		margin-left: 5%;
		text-align: center;
		.image{
			width:50%;
			margin-left:25%;
			box-sizing: border-box;
			height:120px;
			image{
				width:100%;
				height:100%;
				border-radius: 10px;
				box-shadow: 2px 4px 3px #646464;
			}
		}
		
		.title{
			color:#f0f0f0;
			font-size:24px;
			line-height:60px;
		}
		.sub_title{
			margin-top:5px;
			font-size:18px;
			line-height:30px;
			color:#c5c5c5;
		}
		.state{
			margin-top:8px;
			font-size:16px;
			line-height:30px;
			color:#080807;
			margin-bottom:20px;
		}
		.slider{
			margin-top: 30px;
			border:1px solid #ccc;
			background: rgba($color: #fafafa, $alpha: .75);
			box-shadow: 2px 3px 2px rgba(233,233,233,.8);
			border-radius: 12px;
			box-sizing: border-box;
			padding:12px 18px;
			.title{
				
				color:#556;
				font-size: 18px;
				display: block;
				
				line-height:30px;
				text-align: left;
				margin-bottom: 15px;
			}
			.button{
				background:rgba(30,25,48,.8);
				color:#F1F1F1;
				box-shadow:4px 3px 5px #462833;
			}
		}
	}
}
</style>
