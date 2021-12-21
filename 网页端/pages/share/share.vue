<template>
	<view class="container">
		<view class="tips">
			该功能暂时没有开通，抱歉哦。
		</view>
		<view class="footer">
			<view class="icon i1" @click="go(1)">首 页</view>
			<view class="icon i2" @click="go(2)">共 享</view>
			<view class="icon i3" @click="go(3)">我 的</view>
		</view>
	</view>
</template>

<script>
	export default{
		data(){
			return{
				now:2,
			}
		},
		methods:{
			go(id){
				if(id != this.now){
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
		}
	}
</script>

<style lang="scss">
	.container{
		width:100vw;
		min-height:100vh;
		box-sizing: border-box;
		background:rgba($color: #747474, $alpha: 1.0);
		box-sizing: border-box;
		padding-top:30px;
		.tips{
			color:#f0eff4;
			text-align:center;
			font-size:45rpx;
			line-height: 90rpx;
		}
	}
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
		.i2{
			color:#24a141;
		}
	}
</style>
