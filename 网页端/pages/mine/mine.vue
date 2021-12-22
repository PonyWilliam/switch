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
					<view class="getCode">
						<button type="default" @click="getCode">{{codetip}}</button>
					</view>
				</view>
				<view class="registry_tip" @click="go_registry(0)">
					已又账号？点我登录
				</view>
				<view class="submit">
					<button type="default" @click="Registry()">注 册</button>
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
	 function isPoneAvailable(phone) {
	    var myreg=/^[1][3,4,5,6,7,8,9][0-9]{9}$/;
	    if (!myreg.test(phone)) {
	        return false;
	    } else {
			return true;
	    }
	}
	export default{
		data(){
			return{
				username:'',
				password:'',
				now:3,
				registry:0,
				phone:'',
				checkcode:'',
				time:0,
				codetip:'获取验证码',
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
						url:`${configs.default.backurl}/user/login`,
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
								uni.showToast({
									title:res.data.msg,
									icon:"error",
									duration:3000,
								})
								return
							}
							uni.setStorageSync("token",res.data.token)//存储token到storage，后续通过这个作为身份验证
							if(!this.registry){
								uni.showToast({
									title:'登录成功，即将跳转',
									duration:2000,
								})
								setTimeout(()=>{
									uni.redirectTo({
										url:'../index/index'
									})
								},2000)
							}else{
								setTimeout(()=>{
									uni.redirectTo({
										url:'../index/index'
									})
								},1)
							}
						}
					})
				}
			},
			getCode(){
				if(!isPoneAvailable(this.phone)){
					uni.showToast({
						title:'请输入正确的手机号',
						icon:'error',
						duration:2000,
					})
					return
				}
				if(this.time > 0){
					uni.showToast({
						title:'60秒内只能发送一次验证码',
						icon:'error',
						duration:2000,
					})
					return
				}
				//请求发送验证码
				uni.request({
					url:`${configs.default.backurl}/user/apply`,
					data:{
						"phone":this.phone,
					},
					method:"POST",
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},success:res=>{
						console.log(res.data)
						if(res.data.code != 200){
							uni.showToast({
								title:res.data.msg,
								icon:'error',
								duration:2000
							})
							return
						}
						uni.setStorageSync('phone',this.phone)
						uni.showToast({
							title:'验证码发送成功',
							duration:2000
						})
						//启动定时器
						let that = this
						that.time = 60
						let timer = setInterval(()=>{
							if(--that.time <= 0){
								that.codetip = '获取验证码'
								clearInterval(timer)
							}else{
								that.codetip = `${that.time}s`
							}
						},1000)
					}
				})
			},
			Registry(){
				if(this.username.length < 5){
					uni.showToast({
						title:'用户名至少需要5位',
						icon:'error',
						duration:2000
					})
					return
				}
				if(this.password.length < 6){
					uni.showToast({
						title:'密码至少需要5位',
						icon:'error',
						duration:2000
					})
					return
				}
				if(this.phone.length != 11 || this.phone != uni.getStorageSync('phone')){
					uni.showToast({
						title:'手机号错误',
						icon:'error',
						duration:2000
					})
					return
				}
				if(this.checkcode.length != 6){
					uni.showToast({
						title:'验证码错误',
						icon:'error',
						duration:2000
					})
					return
				}
				//请求注册
				uni.request({
					url:`${configs.default.backurl}/user/registry`,
					data:{
						"username":this.username,
						"password":this.password,
						"phone":this.phone,
						"code":this.checkcode,
					},
					method:"POST",
					header:{
						"content-type":"application/x-www-form-urlencoded"
					},
					success:res=>{
						if(res.data.code != 200){
							uni.showToast({
								title:res.data.msg,
								icon:'error',
								duration:2000
							})
						}
						return
						//注册成功
						uni.showToast({
							title:'注册成功',
							duration:1500
						})
						let that = this
						setTimeout(()=>{
							that.login()
						},1500)
					}
				})
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
