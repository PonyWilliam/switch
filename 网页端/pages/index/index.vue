<template>
	<view class="container">
		<view class="header">
			<!-- 预留区域 -->
		</view>
		<view class="content">
			<!-- 展示区域 -->
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
			<!-- 预留区域 -->
		</view>
	</view>
</template>
<style lang="scss">
	.container{
		width:100%;
		min-height:100vh;
		box-sizing: border-box;
		background:rgba($color: #f2f2f2, $alpha: 1.0)
	}
	.cards{
		margin-top:100px;
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
			background:rgba($color: #fff, $alpha: 0.65);
			padding:5px 10px;
			border-radius: 12px;
			box-shadow:2px 3px 4px #d8d8e7;
			.card_image{
				padding:10px 5%;
				width:90%;
				box-sizing: border-box;
				height:120px;
				image{
					width:100%;
					height:100%;
					border-radius: 10px;
					border: .5px solid #ececec;
					box-shadow: 2px 2px 3px #c4c4f2;
				}
			}
			.card_title{
				margin-top:10px;
				font-size:17px;
				line-height:35px;
				color:#333;
			}
			.card_subtitle{
				margin-top:5px;
				font-size:14px;
				line-height:25px;
				color:#525252;
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
</style>
<script>
	const back_url = "https://switchiot.dadiqq.cn"
	export default{
		data(){
			return{
				//返回数据
				devices:[],
			}
		},
		methods:{
			getData:function(){
				
			},
			go_detail:function(device){
				//主要需要传入key，name，nickname
				uni.navigateTo({
					url:"../detail/detail",
					success:res=>{
						res.eventChannel.emit('acceptDataFromOpenerPage', { data: device})
					}
				})
			}
		},
		onLoad:function(){
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