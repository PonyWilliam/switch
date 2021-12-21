<template>
	<view class="container">
		
		<view class="header">
			<!--
				找到好看的图片素材在做
			-->
		</view>
		
		<view class="content">
			<view class="form login" v-show="registry">
				<view class="input_form username">
					<view class="tips">
						账 号：
					</view>
					<view class="input">
						<input type="text" v-model="username" placeholder="请输入账号"/>
					</view>
				</view>
				<view class="input_form password">
					<view class="tips">
						密 码：
					</view>
					<view class="input">
						<input type="password" v-model="password" placeholder="请输入密码"/>
					</view>
				</view>
				<view class="input_form phone">
					<view class="tips">
						手 机 号：
					</view>
					<view class="input">
						<input type="number" v-model="phone" placeholder="请输入手机号"/>
					</view>
				</view>
				<view class="input_form password">
					<view class="tips">
						验 证 码：
					</view>
					<view class="input">
						<input type="number" v-model="checkcode" placeholder="请输入验证码"/>
					</view>
				</view>
				<view class="registry_tip" @click="go_registry(0)">
					已又账号？点我登录
				</view>
				<view class="submit">
					<button type="default" @click="registry()">注 册</button>
				</view>
			</view>
			<view class="form login" v-show="!registry">
				<view class="input_form username">
					<view class="tips">
						账 号：
					</view>
					<view class="input">
						<input type="text" v-model="username" placeholder="请输入账号"/>
					</view>
				</view>
				<view class="input_form password">
					<view class="tips">
						密 码：
					</view>
					<view class="input">
						<input type="password" v-model="password" placeholder="请输入密码"/>
					</view>
					
				</view>
				<view class="registry_tip" @click="go_registry(1)">
					还没有账号?点我注册
				</view>
				<view class="submit">
					<button type="default" @click="login()">登 陆</button>
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

<script>
	const configs = require('../../utils/config.js')
	export default{
		data(){
			return{
				username:'',
				password:'',
				now:3,
				registry:0,
				phone:'',
				checkcode:'',
			}
		},
		methods:{
			onLoad:function(){
			},
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
			},login:function(){
				//登录方法
				if(this.password == '' || this.username == ''){
					if(this.username == ''){
						uni.showModal({
							title:'提示',
							content:'请输入用户名',
							showCancel:false
						})
					}else{
						uni.showModal({
							title:'提示',
							content:'请输入密码',
							showCancel:false
						})
					}
				}else{
					//登录，请求后台接口
					uni.request({
						url:`${configs.default.backurl}/login`,
						data:{
							"username":this.username,
							"password":this.password
						},
						method:"POST",
						header:{
							"content-type":"application/x-www-form-urlencoded"
						},
						success:res=>{
							console.log(res.data)
							if(res.data.code != 200){
								//出错，报错
							}
						}
					})
				}
			},
			go_registry:function(now){
				this.registry = now
			},
		}
	}
</script>

<style lang="scss">
	.container{
		width:100%;
		min-height:100vh;
		box-sizing: border-box;
		background:rgba($color: #747474, $alpha: 1.0);
		.content{
			.form{
				transition: all .5s;
				width:100%;
				background:#dadada;
				padding:30px 0;
				.registry_tip{
					margin-left:5%;
					width:90%;
					padding:20px 0;
					text-align: right;
					color:#333;
				}
				.input_form{
					display: flex;
					flex-direction: row;
					flex-wrap: nowrap;
					justify-content: center;
					padding: 10px 20px;
					.tips{
						border-right: 2px solid #ececec;
					}
					view{
						display: inline-flex;
						box-sizing: border-box;
						line-height: 80rpx;
						font-size:28rpx;
						input{
							margin-left:12px;
							border-radius: 12rpx;
							padding:6px 10px;
							width:100%;
							height:100%;
							box-sizing: border-box;
							background: rgba($color: #e1e1e1, $alpha: .9);
						}
					}
				}
				.submit{
					margin-top:20px;
					display: flex;
					width:100%;
					justify-content: center;
					button{
						background:#333333;
						color:skyblue;
						border-radius: 15px;
					}
				}
				
			}
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
		.i3{
			color:#24a141;
		}
	}
</style>
