package project

import "encoding/base64"

var buildinTemplate = map[string]string{}

func init() {
	buildinTemplate["main.tpl"] = `cGFja2FnZSBtYWluCgppbXBvcnQgKAoJImdpdGh1Yi5jb20veGZhbGkvbmV2ZS1jb3JlIgoJImdpdGh1Yi5jb20veGZhbGkvbmV2ZS1jb3JlL2JlYW4iCgkiZ2l0aHViLmNvbS94ZmFsaS9uZXZlLWNvcmUvYm9vdCIKCSJnaXRodWIuY29tL3hmYWxpL25ldmUtY29yZS9wcm9jZXNzb3IiCgkiZ2l0aHViLmNvbS94ZmFsaS9uZXZlLWxvZ2dlci94bG9nbmV2ZSIKCSJnaXRodWIuY29tL3hmYWxpL25ldmUtdXRpbHMvbmV2ZXJyb3IiCgl7ey0gaWYgbm90IC5Ob1dlYn19CgkiZ2l0aHViLmNvbS94ZmFsaS9uZXZlLXdlYiIKCXt7LSBlbmR9fQoJe3stIGlmIG5vdCAuTm9EYXRhYmFzZX19CgkiZ2l0aHViLmNvbS94ZmFsaS9uZXZlLWRhdGFiYXNlL2dvYmF0aXNldmUiCgl7ey0gZW5kfX0KKQoKe3tpZiBub3QgLk5vU3dhZ2dlcn19Ci8vIGF1dGhvcgovLyBAdGl0bGUgQXdlc29tZSBwcm9qZWN0Ci8vIEB2ZXJzaW9uIHYxLjAuMAovLyBAZGVzY3JpcHRpb24gQXdlc29tZSBwcm9qZWN0CgovLyBAY29udGFjdC5uYW1lIGF1dGhvcgovLyBAY29udGFjdC5lbWFpbCBhdXRob3JAbWFpbC5vcmcKe3tlbmR9fQpmdW5jIG1haW4oKSB7CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXIoeGxvZ25ldmUuTmV3TG9nZ2VyUHJvY2Vzc29yKCkpKQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyKHByb2Nlc3Nvci5OZXdWYWx1ZVByb2Nlc3NvcigpKSkKCXt7LSBpZiBub3QgLk5vV2VifX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlcihnaW5ldmUuTmV3UHJvY2Vzc29yKCkpKQoJe3stIGVuZH19Cgl7ey0gaWYgbm90IC5Ob0RhdGFiYXNlfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3Rlcihnb2JhdGlzZXZlLk5ld1Byb2Nlc3NvcigpKSkKCXt7LSBlbmR9fQoJYm9vdC5SdW4oKQp9`
	buildinTemplate["mod.tpl"] = `bW9kdWxlIHt7Lk1vZHVsZX19CgpnbyB7ey5Hb1ZlcnNpb259fQoKcmVxdWlyZSAoCglnaXRodWIuY29tL3hmYWxpL25ldmUtY29yZSB2MC4yLjggLy8gaW5kaXJlY3QKCWdpdGh1Yi5jb20veGZhbGkveGxvZyB2MC4xLjYgLy8gaW5kaXJlY3QKCWdpdGh1Yi5jb20veGZhbGkvcmVmbGVjdGlvbiB2MC4wLjAtMjAyMjA3MDUxMzU1MzEtNDY0YmEzMjAxNjcxIC8vIGluZGlyZWN0CglnaXRodWIuY29tL3hmYWxpL25ldmUtbG9nZ2VyIHYwLjAuMC0yMDIyMDUyNDE1MTI1OC04NWJjOTQwMzJkMjggLy8gaW5kaXJlY3QKCWdpdGh1Yi5jb20veGZhbGkvbmV2ZS11dGlscyB2MC4wLjEgLy8gaW5kaXJlY3QKCXt7LSBpZiBub3QgLk5vV2VifX0KCWdpdGh1Yi5jb20vZ2luLWNvbnRyaWIvc3NlIHYwLjEuMCAvLyBpbmRpcmVjdAoJZ2l0aHViLmNvbS9naW4tZ29uaWMvZ2luIHYxLjYuMyAvLyBpbmRpcmVjdAoJZ2l0aHViLmNvbS94ZmFsaS9uZXZlLXdlYiB2MC4wLjkgLy8gaW5kaXJlY3QKCXt7LSBlbmR9fQoJe3stIGlmIG5vdCAuTm9Td2FnZ2VyfX0KCWdpdGh1Yi5jb20vc3dhZ2dvL2dpbi1zd2FnZ2VyIHYxLjMuMCAvLyBpbmRpcmVjdAoJZ2l0aHViLmNvbS9zd2FnZ28vc3dhZyB2MS41LjEgLy8gaW5kaXJlY3QKCXt7LSBlbmR9fQoJe3stIGlmIG5vdCAuTm9EYXRhYmFzZX19CglnaXRodWIuY29tL3hmYWxpL25ldmUtZGF0YWJhc2UgdjAuMC40IC8vIGluZGlyZWN0CglnaXRodWIuY29tL3hmYWxpL2dvYmF0aXMgdjAuMi42IC8vIGluZGlyZWN0CglnaXRodWIuY29tL3hmYWxpL3BhZ2VoZWxwZXIgdjAuMi4xIC8vIGluZGlyZWN0Cgl7ey0gZW5kfX0KKQo=`
}

func getBuildTemplate(name string) string {
	d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	if err != nil {
		return ""
	}
	return string(d)
}