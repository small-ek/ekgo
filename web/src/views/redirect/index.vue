<script>
	export default {
		created() {
			const {
				params,
				query
			} = this.$route
			const {
				path
			} = params
			if (path == undefined) {
				this.$router.replace({
					path: '/',
					query
				})
			} else {
				this.$router.replace({
					path: '/' + path,
					query
				})
			}

		},
		render: function(h) {
			return h() // avoid warning message
		}
	}
</script>
