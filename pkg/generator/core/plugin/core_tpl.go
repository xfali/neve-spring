package plugin

import "encoding/base64"

var buildinTemplate = map[string]string{}

func init() {
	buildinTemplate["core.tmpl"] = `e3stICRJbnN0YW5jZVZhciA6PSBjb25jYXQgIl9wcm94eSIgLk5hbWUgIkluc3RhbmNlIiAtfX0Ke3stICRUeXBlTmFtZSA6PSAuVHlwZU5hbWUgLX19Cnt7aWYgLkNvbnRyb2xsZXJNYXJrZXJ9fQoJe3skSW5zdGFuY2VWYXJ9fSA6PSAme3suVHlwZU5hbWV9fXt9Cnt7LSBpZiAuQ29udHJvbGxlck1hcmtlci5WYWx1ZX19CgkvLyByZWdpc3RlciB7ey5UeXBlTmFtZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5Db250cm9sbGVyTWFya2VyLlZhbHVlfX0iLCB7eyRJbnN0YW5jZVZhcn19KSkKe3stIGVsc2V9fQoJLy8gcmVnaXN0ZXIge3suVHlwZU5hbWV9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbih7eyRJbnN0YW5jZVZhcn19KSkKe3stIGVuZH19Cnt7ZW5kfX0Ke3tpZiAuU2VydmljZU1hcmtlcn19Cgl7eyRJbnN0YW5jZVZhcn19IDo9ICZ7ey5UeXBlTmFtZX19e30Ke3stIGlmIC5TZXJ2aWNlTWFya2VyLlZhbHVlfX0KCS8vIHJlZ2lzdGVyIHt7LlR5cGVOYW1lfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW5CeU5hbWUoInt7LlNlcnZpY2VNYXJrZXIuVmFsdWV9fSIsIHt7JEluc3RhbmNlVmFyfX0pKQp7ey0gZWxzZX19CgkvLyByZWdpc3RlciB7ey5UeXBlTmFtZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKHt7JEluc3RhbmNlVmFyfX0pKQp7ey0gZW5kfX0Ke3tlbmR9fQp7e2lmIC5Db21wb25lbnRNYXJrZXJ9fQoJe3skSW5zdGFuY2VWYXJ9fSA6PSAme3suVHlwZU5hbWV9fXt9Cnt7LSBpZiAuQ29tcG9uZW50TWFya2VyLlZhbHVlIH19CgkvLyByZWdpc3RlciB7ey5UeXBlTmFtZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5Db21wb25lbnRNYXJrZXIuVmFsdWV9fSIsIHt7JEluc3RhbmNlVmFyfX0pKQp7ey0gZWxzZX19CgkvLyByZWdpc3RlciB7ey5UeXBlTmFtZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKHt7JEluc3RhbmNlVmFyfX0pKQp7ey0gZW5kfX0Ke3tlbmR9fQp7e2lmIC5CZWFuTWFya2VyfX0Ke3tpZiAuQmVhbk1hcmtlci5WYWx1ZX19CgkvLyByZWdpc3RlciB7ey5UeXBlTmFtZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5CZWFuTWFya2VyLlZhbHVlfX0iLCB7ey5UeXBlTmFtZX19KSkKe3tlbHNlfX0KCS8vIHJlZ2lzdGVyIHt7LlR5cGVOYW1lfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oe3suVHlwZU5hbWV9fSkpCnt7ZW5kfX0Ke3tlbmR9fQoKe3stIHJhbmdlIC5NZXRob2RzIH19Cnt7aWYgLkJlYW5NYXJrZXIgfX0Ke3tpZiAuQmVhbk1hcmtlci5WYWx1ZX19CgkvLyByZWdpc3RlciB7eyRUeXBlTmFtZX19Lnt7Lk5hbWV9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbkJ5TmFtZSgie3suQmVhbk1hcmtlci5WYWx1ZX19Iiwge3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0pKQp7e2Vsc2V9fQoJLy8gcmVnaXN0ZXIge3skVHlwZU5hbWV9fS57ey5OYW1lfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oe3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0pKQp7e2VuZH19Cnt7ZW5kfX0Ke3stIGVuZH19`
}

func getBuildTemplate(name string) string {
	d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	if err != nil {
		return ""
	}
	return string(d)
}