package router

import (
	"fmt"
	"golang-kit/net/context"
)

func jsonp(c context.Context, bs []byte) []byte {
	req := c.Request()
	resp := c.Response()
	params := req.Form
	if params.Get("jsonp") == "jsonp" {
		if cb := params.Get("callback"); cb != "" {
			bs = []byte(fmt.Sprintf("%s(%s)", cb, bs))
		}
		if script := params.Get("script"); script == "script" {
			resp.Header().Set("Content-Type", "text/html;charset=utf-8")
			bs = []byte(fmt.Sprintf(
				`<script type="text/javascript">
					document.domain = 'xxx.com';
					window.parent.%s;
				</script>`, bs))
		}
	}
	return bs
}
