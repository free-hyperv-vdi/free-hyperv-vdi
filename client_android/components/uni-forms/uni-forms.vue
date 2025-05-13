<template>
	<!--   -->
	<view class="uni-forms" :class="{'uni-forms--top':!border}">
		<form @submit.stop="submitForm" @reset="resetForm">
			<slot></slot>
		</form>
	</view>
</template>

<script>
	
	import Vue from 'vue'
	Vue.prototype.binddata = function(name, value, formName) {
		if (formName) {
			this.$refs[formName].setValue(name, value)
		} else {
			let formVm
			for (let i in this.$refs) {
				const vm = this.$refs[i]
				if (vm && vm.$options && vm.$options.name === 'uniForms') {
					formVm = vm
					break
				}
			}
			if (!formVm) return console.error('当前 uni-froms 组件缺少 ref 属性')
			formVm.setValue(name, value)
		}
	}

	import Validator from './validate.js'

	export default {
		name: 'uniForms',
		props: {
			value: {
				type: Object,
				default () {
					return {}
				}
			},
			
			rules: {
				type: Object,
				default () {
					return {}
				}
			},
			
			validateTrigger: {
				type: String,
				default: ''
			},
			
			labelPosition: {
				type: String,
				default: 'left'
			},
			
			labelWidth: {
				type: [String, Number],
				default: 65
			},
			
			labelAlign: {
				type: String,
				default: 'left'
			},
			errShowType: {
				type: String,
				default: 'undertext'
			},
			border: {
				type: Boolean,
				default: false
			}
		},
		data() {
			return {
				formData: {}
			};
		},
		watch: {
			rules(newVal) {
				this.init(newVal)
			},
			trigger(trigger) {
				this.formTrigger = trigger
			},
			value: {
				handler(newVal) {
					if (this.isChildEdit) {
						this.isChildEdit = false
						return
					}
					this.childrens.forEach((item) => {
						if (item.name) {
							const formDataValue = newVal.hasOwnProperty(item.name) ? newVal[item.name] : null
							this.formData[item.name] = this._getValue(item, formDataValue)
						}
					})
				},
				deep: true
			}
		},
		created() {
			let _this = this
			this.childrens = []
			this.inputChildrens = []
			this.checkboxChildrens = []
			this.formRules = []
			this.init(this.rules)
		},
		methods: {
			init(formRules) {
				if (Object.keys(formRules).length > 0) {
					this.formTrigger = this.trigger
					this.formRules = formRules
					if (!this.validator) {
						this.validator = new Validator(formRules)
					}
				}
				this.childrens.forEach((item) => {
					item.init()
				})
			},
			
			setRules(formRules) {
				this.init(formRules)
			},
			
			setValue(name, value, callback) {
				let example = this.childrens.find(child => child.name === name)
				if (!example) return null
				this.isChildEdit = true
				value = this._getValue(example, value)
				this.formData[name] = value
				example.val = value
				this.$emit('input', Object.assign({}, this.value, this.formData))
				return example.triggerCheck(value, callback)
			},

			
			submitForm(event) {
				const value = event.detail.value
				return this.validateAll(value || this.formData, 'submit')
			},
			
			resetForm(event) {
				this.childrens.forEach(item => {
					item.errMsg = ''
					const inputComp = this.inputChildrens.find(child => child.rename === item.name)
					if (inputComp) {
						inputComp.errMsg = ''
						inputComp.$emit('input', inputComp.multiple?[]:'')
					}
				})

				this.isChildEdit = true
				this.childrens.forEach((item) => {
					if (item.name) {
						this.formData[item.name] = this._getValue(item, '')
					}
				})

				this.$emit('input', this.formData)
				this.$emit('reset', event)
			},

			
			validateCheck(validate) {
				if (validate === null) validate = null
				this.$emit('validate', validate)
			},
			
			async validateAll(invalidFields, type, callback) {

				this.childrens.forEach(item => {
					item.errMsg = ''
				})

				let promise;
				if (!callback && typeof callback !== 'function' && Promise) {
					promise = new Promise((resolve, reject) => {
						callback = function(valid, invalidFields) {
							!valid ? resolve(invalidFields) : reject(valid);
						};
					});
				}

				let fieldsValue = {}
				let tempInvalidFields = Object.assign({}, invalidFields)

				Object.keys(this.formRules).forEach(item => {
					const values = this.formRules[item]
					const rules = (values && values.rules) || []
					let isNoField = false
					for (let i = 0; i < rules.length; i++) {
						const rule = rules[i]
						if (rule.required) {
							isNoField = true
							break
						}
					}

					
					if (!isNoField && (!tempInvalidFields[item] && tempInvalidFields[item] !== false)) {
						delete tempInvalidFields[item]
					}

				})
				
				for (let i in this.formRules) {
					for (let j in tempInvalidFields) {
						if (i === j) {
							fieldsValue[i] = tempInvalidFields[i]
						}
					}
				}
				let result = []
				let example = null
				if (this.validator) {
					for (let i in fieldsValue) {
						const resultData = await this.validator.validateUpdate({
							[i]: fieldsValue[i]
						}, this.formData)
						if (resultData) {
							example = this.childrens.find(child => child.name === resultData.key)
							const inputComp = this.inputChildrens.find(child => child.rename === example.name)
							if (inputComp) {
								inputComp.errMsg = resultData.errorMessage
							}
							result.push(resultData)
							if (this.errShowType === 'undertext') {
								if (example) example.errMsg = resultData.errorMessage
							} else {
								if (this.errShowType === 'toast') {
									uni.showToast({
										title: resultData.errorMessage || '校验错误',
										icon: 'none'
									})
									break
								} else if (this.errShowType === 'modal') {
									uni.showModal({
										title: '提示',
										content: resultData.errorMessage || '校验错误'
									})
									break
								} else {
									if (example) example.errMsg = resultData.errorMessage
								}
							}
						}
					}
				}

				if (Array.isArray(result)) {
					if (result.length === 0) result = null
				}
				if (type === 'submit') {
					this.$emit('submit', {
						detail: {
							value: invalidFields,
							errors: result
						}
					})
				} else {
					this.$emit('validate', result)
				}
				callback && typeof callback === 'function' && callback(result, invalidFields)
				if (promise && callback) {
					return promise
				} else {
					return null
				}
			},

			
			submit(callback) {
				
				return this.validateAll(this.formData, 'submit', callback)
			},

			
			validate(callback) {
				return this.validateAll(this.formData, '', callback)
			},

			
			validateField(props, callback) {
				props = [].concat(props);
				let invalidFields = {}
				this.childrens.forEach(item => {
					
					if (props.indexOf(item.name) !== -1) {
						invalidFields = Object.assign({}, invalidFields, {
							[item.name]: this.formData[item.name]
						})
					}
					

				})
				return this.validateAll(invalidFields, '', callback)
			},

			
			resetFields() {
				this.resetForm()
			},

			
			clearValidate(props) {
				props = [].concat(props);
				this.childrens.forEach(item => {
					const inputComp = this.inputChildrens.find(child => child.rename === item.name)
					if (props.length === 0) {
						item.errMsg = ''
						if (inputComp) {
							inputComp.errMsg = ''
						}
					} else {
						if (props.indexOf(item.name) !== -1) {
							item.errMsg = ''
							if (inputComp) {
								inputComp.errMsg = ''
							}
						}
					}
				})
			},
			
			_getValue(item, value) {
				const rules = item.formRules.rules || []
				const isRuleNum = rules.find(val => val.format && this.type_filter(val.format))
				const isRuleBool = rules.find(val => val.format && val.format === 'boolean' || val.format === 'bool')
				
				if (isRuleNum) {
					value = value === '' || value === null ? null : Number(value)
				}
				
				if (isRuleBool) {
					value = !value ? false : true
				}
				return value
			},
			
			type_filter(format) {
				return format === 'int' || format === 'double' || format === 'number'
			}
		}
	}
</script>

<style lang="scss" scoped>
	.uni-forms {
		overflow: hidden;
		
		
	}

	.uni-forms--top {
		padding: 10px 15px;
		
	}
</style>
