<template>
	<view class="uni-load-more" @click="onClick">
		<!-- #ifdef APP-NVUE -->
		<loading-indicator v-if="!webviewHide && status === 'loading' && showIcon" :style="{color: color,width:iconSize+'px',height:iconSize+'px'}" :animating="true" class="uni-load-more__img uni-load-more__img--nvue"></loading-indicator>
		<!-- #endif -->
		<!-- #ifdef H5 -->
		<svg width="24" height="24" viewBox="25 25 50 50" v-if="!webviewHide && (iconType==='circle' || iconType==='auto' && platform === 'android') && status === 'loading' && showIcon"
		:style="{width:iconSize+'px',height:iconSize+'px'}" class="uni-load-more__img uni-load-more__img--android-H5">
			<circle cx="50" cy="50" r="20" fill="none" :style="{color:color}" :stroke-width="3"></circle>
		</svg>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE || H5 -->
		<view v-if="!webviewHide && (iconType==='circle' || iconType==='auto' && platform === 'android') && status === 'loading' && showIcon"
		:style="{width:iconSize+'px',height:iconSize+'px'}" class="uni-load-more__img uni-load-more__img--android-MP">
			<view :style="{borderTopColor:color,borderTopWidth:iconSize/12}"></view>
			<view :style="{borderTopColor:color,borderTopWidth:iconSize/12}"></view>
			<view :style="{borderTopColor:color,borderTopWidth:iconSize/12}"></view>
		</view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<view v-else-if="!webviewHide && status === 'loading' && showIcon" :style="{width:iconSize+'px',height:iconSize+'px'}" class="uni-load-more__img uni-load-more__img--ios-H5">
			<image src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyJpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMy1jMDExIDY2LjE0NTY2MSwgMjAxMi8wMi8wNi0xNDo1NjoyNyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNiAoV2luZG93cykiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6QzlBMzU3OTlEOUM0MTFFOUI0NTZDNERBQURBQzI4RkUiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6QzlBMzU3OUFEOUM0MTFFOUI0NTZDNERBQURBQzI4RkUiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDpDOUEzNTc5N0Q5QzQxMUU5QjQ1NkM0REFBREFDMjhGRSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDpDOUEzNTc5OEQ5QzQxMUU5QjQ1NkM0REFBREFDMjhGRSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/Pt+ALSwAAA6CSURBVHja1FsLkFZVHb98LM+F5bHL8khA1iSeiyQBCRM+YGqKUnnJTDLGI0BGZlKDIU2MMglUiDApEZvSsZnQtBRJtKwQNKQMFYeRDR10WOLd8ljYXdh+v8v5fR3Od+797t1dnOnO/Ofce77z+J
						 mode="widthFix"></image>
		</view>
		<!-- #endif -->
		<text class="uni-load-more__text" :style="{color: color}">{{ status === 'more' ? contentText.contentdown : status === 'loading' ? contentText.contentrefresh : contentText.contentnomore }}</text>
	</view>
</template>

<script>
	const platform = uni.getSystemInfoSync().platform

	
	export default {
		name: 'UniLoadMore',
		props: {
			status: {
				
				type: String,
				default: 'more'
			},
			showIcon: {
				type: Boolean,
				default: true
			},
			iconType: {
				type: String,
				default: 'auto'
			},
			iconSize: {
				type: Number,
				default: 24
			},
			color: {
				type: String,
				default: '#777777'
			},
			contentText: {
				type: Object,
				default () {
					return {
						contentdown: '上拉显示更多',
						contentrefresh: '正在加载...',
						contentnomore: '没有更多数据了'
					}
				}
			}
		},
		data() {
			return {
                webviewHide: false,
				platform: platform
			}
		},
		
		computed:{
			iconSnowWidth(){
				return (Math.floor(this.iconSize/24)||1)*2
			}
		},
		
		mounted() {
			
			var pages = getCurrentPages();
			var page = pages[pages.length - 1];
			var currentWebview = page.$getAppWebview();
			currentWebview.addEventListener('hide', () => {
				this.webviewHide = true
			})
			currentWebview.addEventListener('show', () => {
				this.webviewHide = false
			})
			
		},
		methods: {
			onClick() {
				this.$emit('clickLoadMore', {
					detail: {
						status: this.status,
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>

	.uni-load-more {
		
		display: flex;
		
		flex-direction: row;
		height: 40px;
		align-items: center;
		justify-content: center;
	}

	.uni-load-more__text {
		font-size: 15px;
	}

	.uni-load-more__img {
		width: 24px;
		height: 24px;
		margin-right: 8px;
	}

	.uni-load-more__img--nvue {
		color: #666666;
	}

	.uni-load-more__img--android,
	.uni-load-more__img--ios {
		width: 24px;
		height: 24px;
		transform: rotate(0deg);
	}

	
	.uni-load-more__img--android {
		animation: loading-ios 1s 0s linear infinite;
	}

	@keyframes loading-android {
		0% {
			transform: rotate(0deg);
		}

		100% {
			transform: rotate(360deg);
		}
	}

	.uni-load-more__img--ios-H5 {
		position: relative;
		animation: loading-ios-H5 1s 0s step-end infinite;
	}

	.uni-load-more__img--ios-H5>image {
		position: absolute;
		width: 100%;
		height: 100%;
		left: 0;
		top: 0;
	}

	@keyframes loading-ios-H5 {
		0% {
			transform: rotate(0deg);
		}

		8% {
			transform: rotate(30deg);
		}

		16% {
			transform: rotate(60deg);
		}

		24% {
			transform: rotate(90deg);
		}

		32% {
			transform: rotate(120deg);
		}

		40% {
			transform: rotate(150deg);
		}

		48% {
			transform: rotate(180deg);
		}

		56% {
			transform: rotate(210deg);
		}

		64% {
			transform: rotate(240deg);
		}

		73% {
			transform: rotate(270deg);
		}

		82% {
			transform: rotate(300deg);
		}

		91% {
			transform: rotate(330deg);
		}

		100% {
			transform: rotate(360deg);
		}
	}

	

	
	.uni-load-more__img--android-H5 {
		animation: loading-android-H5-rotate 2s linear infinite;
		transform-origin: center center;
	}

	.uni-load-more__img--android-H5>circle {
		display: inline-block;
		animation: loading-android-H5-dash 1.5s ease-in-out infinite;
		stroke: currentColor;
		stroke-linecap: round;
	}

	@keyframes loading-android-H5-rotate {
		0% {
			transform: rotate(0deg);
		}

		100% {
			transform: rotate(360deg);
		}
	}

	@keyframes loading-android-H5-dash {
		0% {
			stroke-dasharray: 1, 200;
			stroke-dashoffset: 0;
		}

		50% {
			stroke-dasharray: 90, 150;
			stroke-dashoffset: -40;
		}

		100% {
			stroke-dasharray: 90, 150;
			stroke-dashoffset: -120;
		}
	}

	

	
	.uni-load-more__img--android-MP {
		position: relative;
		width: 24px;
		height: 24px;
		transform: rotate(0deg);
		animation: loading-ios 1s 0s ease infinite;
	}

	.uni-load-more__img--android-MP>view {
		position: absolute;
		box-sizing: border-box;
		width: 100%;
		height: 100%;
		border-radius: 50%;
		border: solid 2px transparent;
		border-top: solid 2px #777777;
		transform-origin: center;
	}

	.uni-load-more__img--android-MP>view:nth-child(1){
		animation: loading-android-MP-1 1s 0s linear infinite;
	}

	.uni-load-more__img--android-MP>view:nth-child(2){
		animation: loading-android-MP-2 1s 0s linear infinite;
	}

	.uni-load-more__img--android-MP>view:nth-child(3){
		animation: loading-android-MP-3 1s 0s linear infinite;
	}

	@keyframes loading-android {
		0% {
			transform: rotate(0deg);
		}

		100% {
			transform: rotate(360deg);
		}
	}

	@keyframes loading-android-MP-1{
		0%{
			transform: rotate(0deg);
		}
		50%{
			transform: rotate(90deg);
		}
		100%{
			transform: rotate(360deg);
		}
	}
	@keyframes loading-android-MP-2{
		0%{
			transform: rotate(0deg);
		}
		50%{
			transform: rotate(180deg);
		}
		100%{
			transform: rotate(360deg);
		}
	}
	@keyframes loading-android-MP-3{
		0%{
			transform: rotate(0deg);
		}
		50%{
			transform: rotate(270deg);
		}
		100%{
			transform: rotate(360deg);
		}
	}
	
</style>
