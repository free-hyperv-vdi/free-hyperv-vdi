<template>
	<view v-if="showPopup" class="uni-popup" :class="[popupstyle, isDesktop ? 'fixforpc-z-index' : '']"
	 @touchmove.stop.prevent="clear">
		<uni-transition v-if="maskShow" class="uni-mask--hook" :mode-class="['fade']" :styles="maskClass" :duration="duration"
		 :show="showTrans" @click="onTap" />
		<uni-transition :mode-class="ani" :styles="transClass" :duration="duration" :show="showTrans" @click="onTap">
			<view class="uni-popup__wrapper-box" @click.stop="clear">
				<slot />
			</view>
		</uni-transition>
		<!-- #ifdef H5 -->
		<keypress v-if="maskShow" @esc="onTap" />
		<!-- #endif -->
	</view>
</template>

<script>
	import popup from './popup.js'
	
	import keypress from './keypress.js'
	
	

	export default {
		name: 'UniPopup',
		components: {
			
			keypress
			
		},
		props: {
			
			animation: {
				type: Boolean,
				default: true
			},
			
			
			type: {
				type: String,
				default: 'center'
			},
			
			maskClick: {
				type: Boolean,
				default: true
			}
		},
		provide() {
			return {
				popup: this
			}
		},
		mixins: [popup],
		watch: {
			
			type: {
				handler: function(newVal) {
					this[this.config[newVal]]()
				},
				immediate: true
			},
			isDesktop: {
				handler: function(newVal) {
					this[this.config[this.type]]()
				},
				immediate: true
			},
			
			maskClick: {
				handler: function(val) {
					this.mkclick = val
				},
				immediate: true
			}
		},
		data() {
			return {
				duration: 300,
				ani: [],
				showPopup: false,
				showTrans: false,
				maskClass: {
					'position': 'fixed',
					'bottom': 0,
					'top': 0,
					'left': 0,
					'right': 0,
					'backgroundColor': 'rgba(0, 0, 0, 0.4)'
				},
				transClass: {
					'position': 'fixed',
					'left': 0,
					'right': 0,
				},
				maskShow: true,
				mkclick: true,
				popupstyle: this.isDesktop ? 'fixforpc-top' : 'top'
			}
		},
		created() {
			this.mkclick = this.maskClick
			if (this.animation) {
				this.duration = 300
			} else {
				this.duration = 0
			}
		},
		methods: {
			clear(e) {
				
				e.stopPropagation()
			},
			open() {
				this.showPopup = true
				this.$nextTick(() => {
					new Promise(resolve => {
						clearTimeout(this.timer)
						this.timer = setTimeout(() => {
							this.showTrans = true
							
							this.$nextTick(() => {
								resolve();
							})
						}, 50);
					}).then(res => {
						
						clearTimeout(this.msgtimer)
						this.msgtimer = setTimeout(() => {
							this.customOpen && this.customOpen()
						}, 100)
						this.$emit('change', {
							show: true,
							type: this.type
						})
					})
				})
			},
			close(type) {
				this.showTrans = false
				this.$nextTick(() => {
					this.$emit('change', {
						show: false,
						type: this.type
					})
					clearTimeout(this.timer)
					
					this.customOpen && this.customClose()
					this.timer = setTimeout(() => {
						this.showPopup = false
					}, 300)
				})
			},
			onTap() {
				if (!this.mkclick) return
				this.close()
			},
			
			top() {
				this.popupstyle = this.isDesktop ? 'fixforpc-top' : 'top'
				this.ani = ['slide-top']
				this.transClass = {
					'position': 'fixed',
					'left': 0,
					'right': 0,
				}
			},
			
			bottom() {
				this.popupstyle = 'bottom'
				this.ani = ['slide-bottom']
				this.transClass = {
					'position': 'fixed',
					'left': 0,
					'right': 0,
					'bottom': 0
				}
			},
			
			center() {
				this.popupstyle = 'center'
				this.ani = ['zoom-out', 'fade']
				this.transClass = {
					'position': 'fixed',
					
					'display': 'flex',
					'flexDirection': 'column',
					
					'bottom': 0,
					'left': 0,
					'right': 0,
					'top': 0,
					'justifyContent': 'center',
					'alignItems': 'center'
				}
			}
		}
	}
</script>
<style lang="scss" scoped>
	.uni-popup {
		position: fixed;
		
		z-index: 99;
		
	}

	.fixforpc-z-index {
		
		z-index: 999;
		
	}

	.uni-popup__mask {
		position: absolute;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		background-color: $uni-bg-color-mask;
		opacity: 0;
	}

	.mask-ani {
		transition-property: opacity;
		transition-duration: 0.2s;
	}

	.uni-top-mask {
		opacity: 1;
	}

	.uni-bottom-mask {
		opacity: 1;
	}

	.uni-center-mask {
		opacity: 1;
	}

	.uni-popup__wrapper {
		
		display: block;
		
		position: absolute;
	}

	.top {
		
		top: var(--window-top);
		
		
		top: 0;
		
	}

	.fixforpc-top {
		top: 0;
	}

	.bottom {
		bottom: 0;
	}

	.uni-popup__wrapper-box {
		
		display: block;
		
		position: relative;
		
		
		padding-bottom: constant(safe-area-inset-bottom);
		padding-bottom: env(safe-area-inset-bottom);
		
	}

	.content-ani {
		
		transition-property: transform, opacity;
		transition-duration: 0.2s;
	}


	.uni-top-content {
		transform: translateY(0);
	}

	.uni-bottom-content {
		transform: translateY(0);
	}

	.uni-center-content {
		transform: scale(1);
		opacity: 1;
	}
</style>
