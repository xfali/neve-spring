package plugin

import "encoding/base64"

var buildinTemplate = map[string]string{}

func init() {
	buildinTemplate["webgin.tmpl"] = `e3stICRUeXBlTmFtZSA6PSAuTmFtZSAtfX0Ke3stICRSb290VXJsIDo9IC5SZXF1ZXN0TWFwcGluZy5WYWx1ZSAtfX0KCnR5cGUgTmV2ZXt7Lk5hbWV9fVByb3h5XyBzdHJ1Y3QgewoJbG9nICB4bG9nLkxvZ2dlcgoJSExvZyBsb2dodHRwLkh0dHBMb2dnZXIgYGluamVjdDoiImAKICAgIHt7Lk5hbWV9fSB7ey5UeXBlTmFtZX19IGBpbmplY3Q6Int7LkNvbnRyb2xsZXIuVmFsdWV9fSJgCn0KCmZ1bmMgTmV3TmV2ZXt7Lk5hbWV9fVByb3h5KCkgKiBOZXZle3suTmFtZX19UHJveHlfIHsKCXJldHVybiAmTmV2ZXt7Lk5hbWV9fVByb3h5X3sKCQlsb2c6IHhsb2cuR2V0TG9nZ2VyKCksCgl9Cn0KCmZ1bmMgKGggKk5ldmV7ey5OYW1lfX1Qcm94eV8pIEh0dHBSb3V0ZXMoZW5naW5lIGdpbi5JUm91dGVyKSB7Cgl7ey0gcmFuZ2UgLk1ldGhvZHMgfX0KCXt7LSBpZiAuUmVxdWVzdE1hcHBpbmcuRmxhZyB9fQoJZW5naW5lLnt7LlJlcXVlc3RNYXBwaW5nLk1ldGhvZH19KCJ7e2NvbmNhdFVybCAkUm9vdFVybCAuUmVxdWVzdE1hcHBpbmcuVmFsdWV9fSIsIGguX3Byb3h5e3suTmFtZX19KQoJe3stIGVuZH19Cgl7ey0gZW5kfX0KfQoKe3stIHJhbmdlIC5NZXRob2RzIH19Cnt7IGlmIC5SZXF1ZXN0TWFwcGluZy5GbGFnIH19CmZ1bmMgKGggKk5ldmV7eyRUeXBlTmFtZX19UHJveHlfKSBfcHJveHl7ey5OYW1lfX0oY3R4ICpnaW4uQ29udGV4dCkgewoJdmFyIHBhcnNlUGFyYW1FcnIgZXJyb3IKe3stIHJhbmdlIC5QYXJhbXN9fQoJe3stIGlmIGVxIC5SZXF1ZXN0VHlwZSAicXVlcnkiIH19Cgl7ey5OYW1lfX1TdHIgOj0gY3R4LlF1ZXJ5KCJ7ey5OYW1lfX0iKQoJaWYge3suTmFtZX19U3RyID09ICIiIHsKCQl7ey0gaWYgLlJlcXVpcmVkfX0KCQlfID0gY3R4LkFib3J0V2l0aEVycm9yKGh0dHAuU3RhdHVzQmFkUmVxdWVzdCwgZm10LkVycm9yZigiUXVlcnkgcGFyYW0ge3suTmFtZX19IGlzIG1pc3NpbmcuICIpKQoJCXJldHVybgoJCXt7LSBlbHNlIC19fQoJCXt7Lk5hbWV9fVN0ciA9ICJ7ey5EZWZhdWx0fX0iCgkJe3stIGVuZH19Cgl9Cgl2YXIge3suTmFtZX19IHt7LlR5cGVOYW1lfX0KCXBhcnNlUGFyYW1FcnIgPSByZWZsZWN0aW9uLlNldFZhbHVlSW50ZXJmYWNlKCZ7ey5OYW1lfX0sIHt7Lk5hbWV9fVN0cikKCWlmIHBhcnNlUGFyYW1FcnIgIT0gbmlsIHsKCQlfID0gY3R4LkFib3J0V2l0aEVycm9yKGh0dHAuU3RhdHVzQmFkUmVxdWVzdCwKCQkJZm10LkVycm9yZigiQ29udmVydCBRdWVyeSBwYXJhbSB7ey5OYW1lfX0gJXMgdG8gdHlwZSB7ey5UeXBlTmFtZX19IGZhaWxlZDogJXYgIiwge3suTmFtZX19U3RyLCBwYXJzZVBhcmFtRXJyKSkKCQlyZXR1cm4KCX0KCXt7LSBlbmQgLX19CgoJe3stIGlmIGVxIC5SZXF1ZXN0VHlwZSAicGF0aCIgfX0KCXt7Lk5hbWV9fVN0ciA6PSBjdHguUGFyYW0oInt7Lk5hbWV9fSIpCglpZiB7ey5OYW1lfX1TdHIgPT0gIiIgewoJCXt7LSBpZiAuUmVxdWlyZWR9fQoJCV8gPSBjdHguQWJvcnRXaXRoRXJyb3IoaHR0cC5TdGF0dXNCYWRSZXF1ZXN0LCBmbXQuRXJyb3JmKCJQYXRoIHBhcmFtIHt7Lk5hbWV9fSBpcyBtaXNzaW5nLiAiKSkKCQlyZXR1cm4KCQl7ey0gZWxzZSAtfX0KCQl7ey5OYW1lfX1TdHIgPSAie3suRGVmYXVsdH19IgoJCXt7LSBlbmR9fQoJfQoJdmFyIHt7Lk5hbWV9fSB7ey5UeXBlTmFtZX19CglwYXJzZVBhcmFtRXJyID0gcmVmbGVjdGlvbi5TZXRWYWx1ZUludGVyZmFjZSgme3suTmFtZX19LCB7ey5OYW1lfX1TdHIpCglpZiBwYXJzZVBhcmFtRXJyICE9IG5pbCB7CgkJXyA9IGN0eC5BYm9ydFdpdGhFcnJvcihodHRwLlN0YXR1c0JhZFJlcXVlc3QsCgkJCWZtdC5FcnJvcmYoIkNvbnZlcnQgUGF0aCBwYXJhbSB7ey5OYW1lfX0gJXMgdG8gdHlwZSB7ey5UeXBlTmFtZX19IGZhaWxlZDogJXYgIiwge3suTmFtZX19U3RyLCBwYXJzZVBhcmFtRXJyKSkKCQlyZXR1cm4KCX0KCXt7LSBlbmQgLX19CgoJe3stIGlmIGVxIC5SZXF1ZXN0VHlwZSAiaGVhZGVyIiB9fQoJe3suTmFtZX19U3RyIDo9IGN0eC5HZXRIZWFkZXIoInt7Lk5hbWV9fSIpCglpZiB7ey5OYW1lfX1TdHIgPT0gIiIgewoJCXt7LSBpZiAuUmVxdWlyZWR9fQoJCV8gPSBjdHguQWJvcnRXaXRoRXJyb3IoaHR0cC5TdGF0dXNCYWRSZXF1ZXN0LCBmbXQuRXJyb3JmKCJIZWFkZXIgcGFyYW0ge3suTmFtZX19IGlzIG1pc3NpbmcuICIpKQogICAgICAgIHJldHVybgoJCXt7LSBlbHNlIC19fQoJCXt7Lk5hbWV9fVN0ciA9ICJ7ey5EZWZhdWx0fX0iCgkJe3stIGVuZH19Cgl9Cgl2YXIge3suTmFtZX19IHt7LlR5cGVOYW1lfX0KCXBhcnNlUGFyYW1FcnIgPSByZWZsZWN0aW9uLlNldFZhbHVlSW50ZXJmYWNlKCZ7ey5OYW1lfX0sIHt7Lk5hbWV9fVN0cikKCWlmIHBhcnNlUGFyYW1FcnIgIT0gbmlsIHsKCQlfID0gY3R4LkFib3J0V2l0aEVycm9yKGh0dHAuU3RhdHVzQmFkUmVxdWVzdCwKCQkJZm10LkVycm9yZigiQ29udmVydCBIZWFkZXIgcGFyYW0ge3suTmFtZX19ICVzIHRvIHR5cGUge3suVHlwZU5hbWV9fSBmYWlsZWQ6ICV2ICIsIHt7Lk5hbWV9fVN0ciwgcGFyc2VQYXJhbUVycikpCgkJcmV0dXJuCgl9Cgl7ey0gZW5kIC19fQoJe3stIGlmIGVxIC5SZXF1ZXN0VHlwZSAiYm9keSIgfX0KCXZhciB7ey5OYW1lfX0ge3suVHlwZU5hbWV9fQoJcGFyc2VQYXJhbUVyciA9IGN0eC5CaW5kKCZ7ey5OYW1lfX0pCglpZiBwYXJzZVBhcmFtRXJyICE9IG5pbCB7CgkJXyA9IGN0eC5BYm9ydFdpdGhFcnJvcihodHRwLlN0YXR1c0JhZFJlcXVlc3QsIHBhcnNlUGFyYW1FcnIpCgkJcmV0dXJuCgl9Cgl7ey0gZW5kIC19fQp7ey0gZW5kfX0KCgl7ey0gJHNpemUgOj0gbGVuIC5QYXJhbXMgfCBhZGQgLTF9fQoJe3stIGlmIC5SZXR1cm5zIH19CglyZXQgOj0gaC57eyRUeXBlTmFtZX19Lnt7Lk5hbWV9fSh7ey0gcmFuZ2UgJGksICR2IDo9IC5QYXJhbXN9fXt7JHYuTmFtZX19e3tpZiBndCAkc2l6ZSAkaX19LCB7e2VuZH19e3tlbmQgLX19KQogICAgY3R4LkpTT04oMjAwLCByZXQpCgl7ey0gZWxzZSAtfX0KCWgue3skVHlwZU5hbWV9fS57ey5OYW1lfX0oe3stIHJhbmdlICRpLCAkdiA6PSAuUGFyYW1zfX17eyR2Lk5hbWV9fXt7aWYgZ3QgJHNpemUgJGl9fSwge3tlbmR9fXt7ZW5kIC19fSkKCXt7LSBlbmR9fQp9Cnt7LSBlbmQgfX0Ke3stIGVuZCB9fQo=`
}

func getBuildTemplate(name string) string {
	d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	if err != nil {
		return ""
	}
	return string(d)
}