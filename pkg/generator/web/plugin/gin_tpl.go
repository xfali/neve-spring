package plugin

import "encoding/base64"

var buildinTemplate = map[string]string{}

func init() {
	buildinTemplate["webgin.tmpl"] = `e3stICRUeXBlTmFtZSA6PSAuTmFtZSAtfX0Ke3stICRSb290VXJsIDo9IC5SZXF1ZXN0TWFwcGluZy5WYWx1ZSAtfX0KCnR5cGUgTmV2ZXt7Lk5hbWV9fVByb3h5XyBzdHJ1Y3QgewoJbG9nICB4bG9nLkxvZ2dlcgoJSExvZyBsb2dodHRwLkh0dHBMb2dnZXIgYGluamVjdDoiImAKICAgIHt7Lk5hbWV9fSAqe3suVHlwZU5hbWV9fSBgaW5qZWN0OiJ7ey5Db250cm9sbGVyLlZhbHVlfX0iYAp9CgpmdW5jIE5ld05ldmV7ey5OYW1lfX1Qcm94eSgpICogTmV2ZXt7Lk5hbWV9fVByb3h5XyB7CglyZXR1cm4gJk5ldmV7ey5OYW1lfX1Qcm94eV97CgkJbG9nOiB4bG9nLkdldExvZ2dlcigpLAoJfQp9CgpmdW5jIChoICpOZXZle3suTmFtZX19UHJveHlfKSBIdHRwUm91dGVzKGVuZ2luZSBnaW4uSVJvdXRlcikgewoJe3stIHJhbmdlIC5NZXRob2RzIH19Cgl7ey0gaWYgLlJlcXVlc3RNYXBwaW5nLkZsYWcgfX0KCWVuZ2luZS57ey5SZXF1ZXN0TWFwcGluZy5NZXRob2R9fSgie3tjb25jYXRVcmwgJFJvb3RVcmwgLlJlcXVlc3RNYXBwaW5nLlZhbHVlfX0ie3tpZiAuTG9nSHR0cE1hcmtlcn19LCB7e2NvbmZpZ0xvZ0h0dHAgLkxvZ0h0dHBNYXJrZXJ9fXt7ZW5kfX0sIGguX3Byb3h5e3suTmFtZX19KQoJe3stIGVuZH19Cgl7ey0gZW5kfX0KfQoKe3stIHJhbmdlIC5NZXRob2RzIH19Cnt7IGlmIC5SZXF1ZXN0TWFwcGluZy5GbGFnIH19Cnt7LSBpZiAuQXBpT3BlcmF0aW9uTWFya2VyfX0KLy8gQFN1bW1hcnkge3suQXBpT3BlcmF0aW9uTWFya2VyLlZhbHVlfX0KLy8gQERlc2NyaXB0aW9uIHt7LkFwaU9wZXJhdGlvbk1hcmtlci5Ob3Rlc319Ci8vIEBUYWdzIHt7LkFwaU9wZXJhdGlvbk1hcmtlci5UYWdzfX0KLy8gQFBhcmFtIHBhZ2UgcXVlcnkgc3RyaW5nIHRydWUgInBhZ2UiCnt7LSByYW5nZSAuUGFyYW1zfX0KCXt7LSBpZiBlcSAuUmVxdWVzdFR5cGUgInF1ZXJ5IiB9fQovLyBAUGFyYW0ge3suTmFtZX19IHF1ZXJ5IHt7LlR5cGVOYW1lfX0ge3suUmVxdWlyZWR9fSAie3suTmFtZX19IgoJe3stIGVuZCAtfX0KCXt7LSBpZiBlcSAuUmVxdWVzdFR5cGUgInBhdGgiIH19Ci8vIEBQYXJhbSB7ey5OYW1lfX0gcGF0aCB7ey5UeXBlTmFtZX19IHt7LlJlcXVpcmVkfX0gInt7Lk5hbWV9fSIKCXt7LSBlbmQgLX19Cgl7ey0gaWYgZXEgLlJlcXVlc3RUeXBlICJoZWFkZXIiIH19Ci8vIEBQYXJhbSB7ey5OYW1lfX0gaGVhZGVyIHt7LlR5cGVOYW1lfX0ge3suUmVxdWlyZWR9fSAie3suTmFtZX19IgoJe3stIGVuZCAtfX0KCXt7LSBpZiBlcSAuUmVxdWVzdFR5cGUgImJvZHkiIH19Ci8vIEBQYXJhbSB7ey5OYW1lfX0gYm9keSB7ey5UeXBlTmFtZX19IHt7LlJlcXVpcmVkfX0gInt7Lk5hbWV9fSIKCXt7LSBlbmQgLX19Cnt7LSBlbmR9fQovLyBAQWNjZXB0IHt7aWYgLlJlcXVlc3RNYXBwaW5nLkNvbnN1bWVzfX17ey5SZXF1ZXN0TWFwcGluZy5Db25zdW1lc319e3tlbHNlfX1qc29ue3tlbmR9fQovLyBAUHJvZHVjZSB7e2lmIC5SZXF1ZXN0TWFwcGluZy5Qcm9kdWNlc319e3suUmVxdWVzdE1hcHBpbmcuUHJvZHVjZXN9fXt7ZWxzZX19anNvbnt7ZW5kfX0Ke3stIGlmIC5SZXR1cm5zIH19Cnt7LSAkUmV0dXJuIDo9IGluZGV4IC5SZXR1cm5zIDB9fQovLyBAU3VjY2VzcyAyMDAge29iamVjdH0ge3skUmV0dXJuLlR5cGVOYW1lfX0Ke3stIGVsc2UgLX19Ci8vIEBTdWNjZXNzIDIwMAp7ey0gZW5kfX0KLy8gQEZhaWx1cmUgNDAwIHtzdHJpbmd9IHN0cmluZyAicGFyYW0gZXJyb3IiCi8vIEBSb3V0ZXIge3tpZiAuQXBpT3BlcmF0aW9uTWFya2VyLlJvdXRlcn19e3suQXBpT3BlcmF0aW9uTWFya2VyLlJvdXRlcn19e3tlbHNlfX17e2NvbmNhdFVybCAkUm9vdFVybCAuUmVxdWVzdE1hcHBpbmcuVmFsdWUgfCBzd2FnZ2VyUm91dGVyfX17e2VuZH19IFt7e3RvTG93ZXIgLlJlcXVlc3RNYXBwaW5nLk1ldGhvZH19XQp7ey0gZW5kfX0KZnVuYyAoaCAqTmV2ZXt7JFR5cGVOYW1lfX1Qcm94eV8pIF9wcm94eXt7Lk5hbWV9fShjdHggKmdpbi5Db250ZXh0KSB7Cgl2YXIgcGFyc2VQYXJhbUVyciBlcnJvcgp7ey0gcmFuZ2UgLlBhcmFtc319Cgl7ey0gaWYgZXEgLlJlcXVlc3RUeXBlICJxdWVyeSIgfX0KCXt7Lk5hbWV9fVN0ciA6PSBjdHguUXVlcnkoInt7Lk5hbWV9fSIpCglpZiB7ey5OYW1lfX1TdHIgPT0gIiIgewoJCXt7LSBpZiAuUmVxdWlyZWR9fQoJCV8gPSBjdHguQWJvcnRXaXRoRXJyb3IoaHR0cC5TdGF0dXNCYWRSZXF1ZXN0LCBmbXQuRXJyb3JmKCJRdWVyeSBwYXJhbSB7ey5OYW1lfX0gaXMgbWlzc2luZy4gIikpCgkJcmV0dXJuCgkJe3stIGVsc2UgLX19CgkJe3suTmFtZX19U3RyID0gInt7LkRlZmF1bHR9fSIKCQl7ey0gZW5kfX0KCX0KCXZhciB7ey5OYW1lfX0ge3suVHlwZU5hbWV9fQoJcGFyc2VQYXJhbUVyciA9IHJlZmxlY3Rpb24uU2V0VmFsdWVJbnRlcmZhY2UoJnt7Lk5hbWV9fSwge3suTmFtZX19U3RyKQoJaWYgcGFyc2VQYXJhbUVyciAhPSBuaWwgewoJCV8gPSBjdHguQWJvcnRXaXRoRXJyb3IoaHR0cC5TdGF0dXNCYWRSZXF1ZXN0LAoJCQlmbXQuRXJyb3JmKCJDb252ZXJ0IFF1ZXJ5IHBhcmFtIHt7Lk5hbWV9fSAlcyB0byB0eXBlIHt7LlR5cGVOYW1lfX0gZmFpbGVkOiAldiAiLCB7ey5OYW1lfX1TdHIsIHBhcnNlUGFyYW1FcnIpKQoJCXJldHVybgoJfQoJe3stIGVuZCAtfX0KCgl7ey0gaWYgZXEgLlJlcXVlc3RUeXBlICJwYXRoIiB9fQoJe3suTmFtZX19U3RyIDo9IGN0eC5QYXJhbSgie3suTmFtZX19IikKCWlmIHt7Lk5hbWV9fVN0ciA9PSAiIiB7CgkJe3stIGlmIC5SZXF1aXJlZH19CgkJXyA9IGN0eC5BYm9ydFdpdGhFcnJvcihodHRwLlN0YXR1c0JhZFJlcXVlc3QsIGZtdC5FcnJvcmYoIlBhdGggcGFyYW0ge3suTmFtZX19IGlzIG1pc3NpbmcuICIpKQoJCXJldHVybgoJCXt7LSBlbHNlIC19fQoJCXt7Lk5hbWV9fVN0ciA9ICJ7ey5EZWZhdWx0fX0iCgkJe3stIGVuZH19Cgl9Cgl2YXIge3suTmFtZX19IHt7LlR5cGVOYW1lfX0KCXBhcnNlUGFyYW1FcnIgPSByZWZsZWN0aW9uLlNldFZhbHVlSW50ZXJmYWNlKCZ7ey5OYW1lfX0sIHt7Lk5hbWV9fVN0cikKCWlmIHBhcnNlUGFyYW1FcnIgIT0gbmlsIHsKCQlfID0gY3R4LkFib3J0V2l0aEVycm9yKGh0dHAuU3RhdHVzQmFkUmVxdWVzdCwKCQkJZm10LkVycm9yZigiQ29udmVydCBQYXRoIHBhcmFtIHt7Lk5hbWV9fSAlcyB0byB0eXBlIHt7LlR5cGVOYW1lfX0gZmFpbGVkOiAldiAiLCB7ey5OYW1lfX1TdHIsIHBhcnNlUGFyYW1FcnIpKQoJCXJldHVybgoJfQoJe3stIGVuZCAtfX0KCgl7ey0gaWYgZXEgLlJlcXVlc3RUeXBlICJoZWFkZXIiIH19Cgl7ey5OYW1lfX1TdHIgOj0gY3R4LkdldEhlYWRlcigie3suTmFtZX19IikKCWlmIHt7Lk5hbWV9fVN0ciA9PSAiIiB7CgkJe3stIGlmIC5SZXF1aXJlZH19CgkJXyA9IGN0eC5BYm9ydFdpdGhFcnJvcihodHRwLlN0YXR1c0JhZFJlcXVlc3QsIGZtdC5FcnJvcmYoIkhlYWRlciBwYXJhbSB7ey5OYW1lfX0gaXMgbWlzc2luZy4gIikpCiAgICAgICAgcmV0dXJuCgkJe3stIGVsc2UgLX19CgkJe3suTmFtZX19U3RyID0gInt7LkRlZmF1bHR9fSIKCQl7ey0gZW5kfX0KCX0KCXZhciB7ey5OYW1lfX0ge3suVHlwZU5hbWV9fQoJcGFyc2VQYXJhbUVyciA9IHJlZmxlY3Rpb24uU2V0VmFsdWVJbnRlcmZhY2UoJnt7Lk5hbWV9fSwge3suTmFtZX19U3RyKQoJaWYgcGFyc2VQYXJhbUVyciAhPSBuaWwgewoJCV8gPSBjdHguQWJvcnRXaXRoRXJyb3IoaHR0cC5TdGF0dXNCYWRSZXF1ZXN0LAoJCQlmbXQuRXJyb3JmKCJDb252ZXJ0IEhlYWRlciBwYXJhbSB7ey5OYW1lfX0gJXMgdG8gdHlwZSB7ey5UeXBlTmFtZX19IGZhaWxlZDogJXYgIiwge3suTmFtZX19U3RyLCBwYXJzZVBhcmFtRXJyKSkKCQlyZXR1cm4KCX0KCXt7LSBlbmQgLX19Cgl7ey0gaWYgZXEgLlJlcXVlc3RUeXBlICJib2R5IiB9fQoJdmFyIHt7Lk5hbWV9fSB7ey5UeXBlTmFtZX19CglwYXJzZVBhcmFtRXJyID0gY3R4LkJpbmQoJnt7Lk5hbWV9fSkKCWlmIHBhcnNlUGFyYW1FcnIgIT0gbmlsIHsKCQlfID0gY3R4LkFib3J0V2l0aEVycm9yKGh0dHAuU3RhdHVzQmFkUmVxdWVzdCwgcGFyc2VQYXJhbUVycikKCQlyZXR1cm4KCX0KCXt7LSBlbmQgLX19Cnt7LSBlbmR9fQoKCXt7LSAkc2l6ZSA6PSBsZW4gLlBhcmFtcyB8IGFkZCAtMX19Cgl7ey0gaWYgLlJldHVybnMgfX0KCXJldCA6PSBoLnt7JFR5cGVOYW1lfX0ue3suTmFtZX19KHt7LSByYW5nZSAkaSwgJHYgOj0gLlBhcmFtc319e3skdi5OYW1lfX17e2lmIGd0ICRzaXplICRpfX0sIHt7ZW5kfX17e2VuZCAtfX0pCiAgICBjdHguSlNPTigyMDAsIHJldCkKCXt7LSBlbHNlIC19fQoJaC57eyRUeXBlTmFtZX19Lnt7Lk5hbWV9fSh7ey0gcmFuZ2UgJGksICR2IDo9IC5QYXJhbXN9fXt7JHYuTmFtZX19e3tpZiBndCAkc2l6ZSAkaX19LCB7e2VuZH19e3tlbmQgLX19KQoJe3stIGVuZH19Cn0Ke3stIGVuZCB9fQp7ey0gZW5kIH19Cg==`
	buildinTemplate["webgin_register.tmpl"] = `CmZ1bmMgaW5pdCgpIHsKe3stIHJhbmdlIC4gfX0Ke3stIGlmIC59fQoJLy8gcmVnaXN0ZXIge3suVHlwZU5hbWV9fSB3ZWIgcHJveHkKCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oTmV3TmV2ZXt7Lk5hbWV9fVByb3h5KCkpKQp7ey0gZW5kfX0Ke3stIGVuZH19Cn0K`
}

func getBuildTemplate(name string) string {
	d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	if err != nil {
		return ""
	}
	return string(d)
}