<template>
	<view class="container">
		<view class="header">
			<!-- 预留区域 -->
			<swiper class="swiper" :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval" :duration="duration">
				<swiper-item>
					<view class="swiper-item">
						<image src="../../static/swiper1.jpg" mode="scaleToFill"></image>
					</view>
				</swiper-item>
				<swiper-item>
					<view class="swiper-item">
						<image src="../../static/swiper2.gif" mode="scaleToFill"></image>
					</view>
				</swiper-item>
				<swiper-item>
					<view class="swiper-item">
						<image src="../../static/swiper3.jpg" mode="scaleToFill"></image>
					</view>
				</swiper-item>
			</swiper>
		</view>
		<view class="content">
			<!-- 展示区域 -->
			<scroll-view class="content_header" scroll-x="true">
				<view v-for="(item,index) in devices_type" :key="index" :class="index==active_text?'active_text normal':'normal'" @click="transform_device(index)">
					{{item}}
				</view>
			</scroll-view>
			<view class="cards">
				<view class="card" v-for="(item, index) in devices" @click="go_detail(item)">
					<view class="card_image"><image src="../../static/switch.jpg" mode="aspectFill"></image></view>
					<view class="card_title">设备:{{item.DeviceName}}</view>
					<view class="card_subtitle">
						<view v-if="item.DeviceStatus == 'ONLINE'" class="cirle green"></view>
						<view class="cirle red" v-else-if="item.DeviceStatus != 'ONLINE' "></view>
						{{item.DeviceStatus == "ONLINE" ? "在线" : "离线"}} | {{item.Nickname}}
					</view>
				</view>
			</view>
		</view>
		<view class="footer">
			<view class="icon i1" @click="go(1)">首 页</view>
			<view class="icon i2" @click="go(2)">共 享</view>
			<view class="icon i3" @click="go(3)">我 的</view>
		</view>
	</view>
</template>
<style lang="scss">
	.footer{
		background:rgba($color: #dfdfdf, $alpha: .65);
		width:80%;
		position: fixed;
		bottom:30px;
		left:10%;
		display: flex;
		height:40px;
		box-sizing: border-box;
		border-radius: 20px;
		backdrop-filter:saturate(180%) blur(20px);
		padding:0 40px;
		flex-wrap: nowrap;
		flex-direction: row;
		justify-content: space-between;
		.icon{
			display: inline-flex;
			line-height:40px;
			color:#333;
		}
		.i1{
			color:#24a141;
		}
	}
	.container{
		width:100%;
		min-height:100vh;
		box-sizing: border-box;
		background:rgba($color: #747474, $alpha: 1.0)
	}
	.header{
		width:90%;
		margin-left:5%;
		box-sizing: border-box;
		padding:15px 0;
		.swiper-item{
			width:100%;
			height:100%;
			border-radius: 24px;
			image{
				box-shadow: 6px 5px 4px rgba(0,25,33,.7),
				-3px 5px 4px rgba(0,25,33,.7),
				3px -5px 4px rgba(0,25,33,.7),
				-3px -5px 4px rgba(0,25,33,.7);
				width:100%;
				height:100%;
				border-radius: 24px;
			}
		}
	}
	.content{
		.content_header{
			margin-left:2.5%;
			padding:18px 20px 18px 0;
			width:95%;
			white-space: nowrap;
			box-sizing: border-box;
			.normal{
				padding:15px 20px;
				background:white;
				margin-right:15px;
				border-radius: 15px;
				display:inline-block;
				text-align: center;
				cursor: pointer;
			}
			.active_text{
				color:#20ffc0;
				background:#2C405A;
				cursor: default;
				transition: all .6s;
			}
		}
		.cards{
				margin-top:10px;
				padding:20px 0;
				width:90%;
				box-sizing: border-box;
				display: flex;
				flex-direction: row;
				margin-left:5%;
				flex-wrap: wrap;
				justify-content: space-between;
				.card{
					width:45%;
					box-sizing: border-box;
					background:rgba($color: #4f4f4f, $alpha: 0.65);
					padding:4px 8px;
					border-radius: 16px;
					box-shadow:2px 3px 4px #101012,
					-2px -3px 4px #101012;
					.card_image{
						padding:10px 5%;
						width:90%;
						box-sizing: border-box;
						height:120px;
						image{
							width:100%;
							height:100%;
							border-radius: 10px;
							border: .5px solid #464646;
							box-shadow: 2px 2px 3px #070709;
						}
					}
					.card_title{
						margin-top:10px;
						font-size:17px;
						line-height:35px;
						color:#c2c2c2;
					}
					.card_subtitle{
						margin-top:5px;
						font-size:14px;
						line-height:25px;
						color:#86e1d4;
						.cirle{
							height:10px;
							width:10px;
							border-radius: 100%;
							display: inline-block;
						}
						.red{
							background:#DD524D;
						}
						.green{
							background:#4CD964;
						}
					}
				}
			}
		
	}
</style>
<script>
	const back_url = "https://switchiot.dadiqq.cn"
	export default{
		data(){
			return{
				//返回数据
				devices:[],
				indicatorDots: true,
				autoplay: true,
				interval: 8000,
				duration: 400,
				devices_type:['智能开关(舵机版)','智能开关(继电器版)'],
				active_text:0,
				now:1,
			}
		},
		methods:{
			getData:function(){
				
			},
			transform_device:function(k){
				console.log(k)
				this.active_text = k
			},
			go_detail:function(device){
				//主要需要传入key，name，nickname
				uni.navigateTo({
					url:"../detail/detail",
					success:res=>{
						res.eventChannel.emit('acceptDataFromOpenerPage', { data: device})
					}
				})
			},
			go(id){
				console.log(id)
				if(id != this.now){
					console.log('不等于')
					if(id == 1){
						uni.redirectTo({
							url:"../index/index"
						})
					}else if(id == 2){
						uni.redirectTo({
							url:"../share/share"
						})
					}else{
						uni.redirectTo({
							url:"../mine/mine"
						})
					}
				}
			}
		},
		onLoad:async function(){
			console.log(uni.getStorageSync('token'))
			let token = uni.getStorageSync('token')
			if(token == ""){
				uni.showToast({
					title:'请先登录',
					duration:1000,
					icon:'error'
				})
				setTimeout(()=>{
					uni.redirectTo({
						url:'../mine/mine'
					})
				},1000)
				return
			}
			const [err,res] = await uni.request({
				url:`${back_url}/user`,
				method:'GET',
				header:{
					"token":token
				},
			})
			if(res.data.code != 200){
				//登录失效
				uni.showToast({
					title:'登录过期，请重新登录',
					duration:1000,
					icon:'error'
				})
				setTimeout(()=>{
					uni.removeStorageSync('token')
					uni.redirectTo({
						url:'../mine/mine'
					})
				},1000)
				return
			}
			uni.showLoading({
				title:'加载数据中···',
			})
			uni.request({
				url:`${back_url}/devices`,
				method:"GET",
				success:res=>{
					this.devices = res.data.data
					uni.hideLoading()
					
				},fail:err=>{
					uni.hideLoading()
					uni.showModal({
						title:"错误",
						content:"无法加载数据，请检查网络连接状态",
						showCancel:false
					})
				}
			})
		},
	}
</script>