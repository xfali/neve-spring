package plugin

import "encoding/base64"

var buildinTemplate = map[string]string{}

func init() {
	buildinTemplate["core.tmpl"] = `e3stICRJbnN0YW5jZVZhciA6PSBjb25jYXQgIl9wcm94eSIgLk5hbWUgIkluc3RhbmNlIiAtfX0Ke3stICRUeXBlTmFtZSA6PSAuVHlwZU5hbWUgLX19Cnt7JFBvc3RDb25zdHJ1Y3RNZXRob2QgOj0gIiJ9fQp7eyRQcmVEZXN0cm95TWV0aG9kIDo9ICIifX0KCnt7LSBpZiAuUG9zdENvbnN0cnVjdE1hcmtlciAtfX0Ke3skUG9zdENvbnN0cnVjdE1ldGhvZCA9IC5Qb3N0Q29uc3RydWN0TWFya2VyLk1ldGhvZE5hbWV9fQp7ey0gZW5kIC19fQoKe3stIGlmIC5QcmVEZXN0cm95TWFya2VyIC19fQp7eyRQcmVEZXN0cm95TWV0aG9kID0gLlByZURlc3Ryb3lNYXJrZXIuTWV0aG9kTmFtZX19Cnt7LSBlbmQgLX19Cgp7ey0gaWYgaGF2ZUJlYW4gLn19CiAgICB7ey0gJEJlYW5OYW1lIDo9IGJlYW5WYWx1ZU9yTmFtZSAuIC19fQogICAge3stIGlmICRCZWFuTmFtZX19CiAgICAgICAgLy8gcmVnaXN0ZXIge3suVHlwZU5hbWV9fQogICAgICAgIHt7LSBpZiBwcm90b3R5cGUgLlNjb3BlTWFya2VyfX0KICAgICAgICAgICAge3stIGlmIGhhdmVBdXRvd2lyZWQgLn19CiAgICAgICAgICAgICAgICB7ey0gJHNpemUgOj0gbGVuIC5GaWVsZHMgfCBhZGQgLTF9fQogICAgICAgICAgICAgICAgbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbkJ5TmFtZSgie3skQmVhbk5hbWV9fSIsIGJlYW4uTmV3Q3VzdG9tQmVhbkZhY3RvcnlXaXRoT3B0cyhmdW5jKHt7cmFuZ2UgJGksICR2IDo9IC5GaWVsZHN9fXt7Zmlyc3RMb3dlciAkdi5OYW1lfX0ge3skdi5UeXBlTmFtZX19e3tpZiBndCAkc2l6ZSAkaX19LCB7e2VuZH19e3tlbmR9fSkgKnt7LlR5cGVOYW1lfX0gewogICAgICAgICAgICAgICAgICAgIHJldCA6PSAme3suVHlwZU5hbWV9fXt9CiAgICAgICAgICAgICAgICAgICAge3stIHJhbmdlIC5GaWVsZHN9fQogICAgICAgICAgICAgICAgICAgIHJldC57ey5OYW1lfX0gPSB7e2ZpcnN0TG93ZXIgLk5hbWV9fQogICAgICAgICAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgICAgICAgICAgICAgIHJldHVybiByZXQKICAgICAgICAgICAgICAgIH0sIGJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLk5hbWVzKFtdc3RyaW5newogICAgICAgICAgICAgICAgICAge3stIHJhbmdlIC5GaWVsZHN9fQogICAgICAgICAgICAgICAgICAgICAgICJ7ey5BdXRvd2lyZWRNYXJrZXIuTmFtZX19e3tpZiBub3QgLkF1dG93aXJlZE1hcmtlci5SZXF1aXJlZH19LG9taXRlcnJvcnt7ZW5kfX0iLAogICAgICAgICAgICAgICAgICAge3stIGVuZH19CiAgICAgICAgICAgICAgICAgICB9KSwKICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUG9zdENvbnN0cnVjdE1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUHJlQWZ0ZXJTZXQoInt7JFBvc3RDb25zdHJ1Y3RNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUHJlRGVzdHJveU1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUG9zdERlc3Ryb3koInt7JFByZURlc3Ryb3lNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICkpKQogICAgICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgICAgICB7ey0gaWYgb3IgJFBvc3RDb25zdHJ1Y3RNZXRob2QgJFByZURlc3Ryb3lNZXRob2R9fQogICAgICAgICAgICAgICAgICAgIG5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW5CeU5hbWUoInt7JEJlYW5OYW1lfX0iLCBiZWFuLk5ld0N1c3RvbUJlYW5GYWN0b3J5V2l0aE9wdHMoZnVuYygpICp7ey5UeXBlTmFtZX19IHsgcmV0dXJuICZ7ey5UeXBlTmFtZX19e30gfSwKICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUG9zdENvbnN0cnVjdE1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUHJlQWZ0ZXJTZXQoInt7JFBvc3RDb25zdHJ1Y3RNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUHJlRGVzdHJveU1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUG9zdERlc3Ryb3koInt7JFByZURlc3Ryb3lNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICkpKQogICAgICAgICAgICAgICAge3stIGVsc2V9fQogICAgICAgICAgICAgICAgICAgIG5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW5CeU5hbWUoInt7JEJlYW5OYW1lfX0iLCBmdW5jKCkgKnt7LlR5cGVOYW1lfX0geyByZXR1cm4gJnt7LlR5cGVOYW1lfX17fSB9KSkKICAgICAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgICAgICB7ey0gZW5kfX0KICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgIHt7JEluc3RhbmNlVmFyfX0gOj0gJnt7LlR5cGVOYW1lfX17fQogICAgICAgICAgICB7ey0gaWYgaGF2ZUF1dG93aXJlZCAufX0KICAgICAgICAgICAgICAgIHt7LSAkc2l6ZSA6PSBsZW4gLkZpZWxkcyB8IGFkZCAtMX19CiAgICAgICAgICAgICAgICBuZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7eyRCZWFuTmFtZX19IiwgYmVhbi5OZXdDdXN0b21CZWFuRmFjdG9yeVdpdGhPcHRzKGZ1bmMoe3tyYW5nZSAkaSwgJHYgOj0gLkZpZWxkc319e3tmaXJzdExvd2VyICR2Lk5hbWV9fSB7eyR2LlR5cGVOYW1lfX17e2lmIGd0ICRzaXplICRpfX0sIHt7ZW5kfX17e2VuZH19KSAqe3suVHlwZU5hbWV9fSB7CiAgICAgICAgICAgICAgICAgICB7ey0gcmFuZ2UgLkZpZWxkc319CiAgICAgICAgICAgICAgICAgICB7eyRJbnN0YW5jZVZhcn19Lnt7Lk5hbWV9fSA9IHt7Zmlyc3RMb3dlciAuTmFtZX19CiAgICAgICAgICAgICAgICAgICB7ey0gZW5kfX0KICAgICAgICAgICAgICAgICAgIHJldHVybiB7eyRJbnN0YW5jZVZhcn19CiAgICAgICAgICAgICAgICB9LCBiZWFuLkN1c3RvbUJlYW5GYWN0b3J5T3B0cy5OYW1lcyhbXXN0cmluZ3sKICAgICAgICAgICAgICAgICAgICAgIHt7LSByYW5nZSAuRmllbGRzfX0KICAgICAgICAgICAgICAgICAgICAgICAgICAie3suQXV0b3dpcmVkTWFya2VyLk5hbWV9fXt7aWYgbm90IC5BdXRvd2lyZWRNYXJrZXIuUmVxdWlyZWR9fSxvbWl0ZXJyb3J7e2VuZH19IiwKICAgICAgICAgICAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgICAgICAgICAgICAgICAgfSksCiAgICAgICAgICAgICAgICAgICAgICB7ey0gaWYgJFBvc3RDb25zdHJ1Y3RNZXRob2R9fWJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLlByZUFmdGVyU2V0KCJ7eyRQb3N0Q29uc3RydWN0TWV0aG9kfX0iKSx7e2VuZH19CiAgICAgICAgICAgICAgICAgICAgICB7ey0gaWYgJFByZURlc3Ryb3lNZXRob2R9fWJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLlBvc3REZXN0cm95KCJ7eyRQcmVEZXN0cm95TWV0aG9kfX0iKSx7e2VuZH19CiAgICAgICAgICAgICAgICAgICAgICkpKQogICAgICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgICAgICB7ey0gaWYgb3IgJFBvc3RDb25zdHJ1Y3RNZXRob2QgJFByZURlc3Ryb3lNZXRob2R9fQogICAgICAgICAgICAgICAgICAgIG5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW5CeU5hbWUoInt7JEJlYW5OYW1lfX0iLCBiZWFuLk5ld0N1c3RvbUJlYW5GYWN0b3J5V2l0aE9wdHMoZnVuYygpICp7ey5UeXBlTmFtZX19IHsgcmV0dXJuIHt7JEluc3RhbmNlVmFyfX0gfSwKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUG9zdENvbnN0cnVjdE1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUHJlQWZ0ZXJTZXQoInt7JFBvc3RDb25zdHJ1Y3RNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUHJlRGVzdHJveU1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUG9zdERlc3Ryb3koInt7JFByZURlc3Ryb3lNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICkpKQogICAgICAgICAgICAgICAge3stIGVsc2V9fQogICAgICAgICAgICAgICAgICAgIG5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW5CeU5hbWUoInt7JEJlYW5OYW1lfX0iLCB7eyRJbnN0YW5jZVZhcn19KSkKICAgICAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgICAgICB7ey0gZW5kfX0KICAgICAgICB7ey0gZW5kfX0KICAgIHt7LSBlbHNlfX0KICAgICAgICAvLyByZWdpc3RlciB7ey5UeXBlTmFtZX19CiAgICAgICAge3stIGlmIHByb3RvdHlwZSAuU2NvcGVNYXJrZXJ9fQogICAgICAgICAgICB7ey0gaWYgaGF2ZUF1dG93aXJlZCAufX0KICAgICAgICAgICAgICAgIHt7LSAkc2l6ZSA6PSBsZW4gLkZpZWxkcyB8IGFkZCAtMX19CiAgICAgICAgICAgICAgICBuZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKGJlYW4uTmV3Q3VzdG9tQmVhbkZhY3RvcnlXaXRoT3B0cyhmdW5jKHt7cmFuZ2UgJGksICR2IDo9IC5GaWVsZHN9fXt7Zmlyc3RMb3dlciAkdi5OYW1lfX0ge3skdi5UeXBlTmFtZX19e3tpZiBndCAkc2l6ZSAkaX19LCB7e2VuZH19e3tlbmR9fSkgKnt7LlR5cGVOYW1lfX0gewogICAgICAgICAgICAgICAgICAgICAgICAgIHJldCA6PSAme3suVHlwZU5hbWV9fXt9CiAgICAgICAgICAgICAgICAgICAgICAgICAge3stIHJhbmdlIC5GaWVsZHN9fQogICAgICAgICAgICAgICAgICAgICAgICAgIHJldC57ey5OYW1lfX0gPSB7e2ZpcnN0TG93ZXIgLk5hbWV9fQogICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgICAgICAgICAgICAgICAgICAgIHJldHVybiByZXQKICAgICAgICAgICAgICAgICAgICAgIH0sIGJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLk5hbWVzKFtdc3RyaW5newogICAgICAgICAgICAgICAgICAgICAgICAge3stIHJhbmdlIC5GaWVsZHN9fQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ7ey5BdXRvd2lyZWRNYXJrZXIuTmFtZX19e3tpZiBub3QgLkF1dG93aXJlZE1hcmtlci5SZXF1aXJlZH19LG9taXRlcnJvcnt7ZW5kfX0iLAogICAgICAgICAgICAgICAgICAgICAgICAge3stIGVuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICB9KSwKICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUG9zdENvbnN0cnVjdE1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUHJlQWZ0ZXJTZXQoInt7JFBvc3RDb25zdHJ1Y3RNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUHJlRGVzdHJveU1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUG9zdERlc3Ryb3koInt7JFByZURlc3Ryb3lNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICAgICAgICkpKQogICAgICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgICAgICB7ey0gaWYgb3IgJFBvc3RDb25zdHJ1Y3RNZXRob2QgJFByZURlc3Ryb3lNZXRob2R9fQogICAgICAgICAgICAgICAgICAgIG5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oYmVhbi5OZXdDdXN0b21CZWFuRmFjdG9yeVdpdGhPcHRzKGZ1bmMoKSAqe3suVHlwZU5hbWV9fSB7IHJldHVybiAme3suVHlwZU5hbWV9fXt9IH0sCiAgICAgICAgICAgICAgICAgICAgICAgICB7ey0gaWYgJFBvc3RDb25zdHJ1Y3RNZXRob2R9fWJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLlByZUFmdGVyU2V0KCJ7eyRQb3N0Q29uc3RydWN0TWV0aG9kfX0iKSx7e2VuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICB7ey0gaWYgJFByZURlc3Ryb3lNZXRob2R9fWJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLlBvc3REZXN0cm95KCJ7eyRQcmVEZXN0cm95TWV0aG9kfX0iKSx7e2VuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICAgKSkpCiAgICAgICAgICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgICAgICAgICAgbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbihmdW5jKCkgKnt7LlR5cGVOYW1lfX0geyByZXR1cm4gJnt7LlR5cGVOYW1lfX17fSB9KSkKICAgICAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgICAgICB7ey0gZW5kfX0KICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgIHt7JEluc3RhbmNlVmFyfX0gOj0gJnt7LlR5cGVOYW1lfX17fQogICAgICAgICAgICB7ey0gaWYgaGF2ZUF1dG93aXJlZCAufX0KICAgICAgICAgICAgICAgIHt7LSAkc2l6ZSA6PSBsZW4gLkZpZWxkcyB8IGFkZCAtMX19CiAgICAgICAgICAgICAgICBuZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKGJlYW4uTmV3Q3VzdG9tQmVhbkZhY3RvcnlXaXRoT3B0cyhmdW5jKHt7cmFuZ2UgJGksICR2IDo9IC5GaWVsZHN9fXt7Zmlyc3RMb3dlciAkdi5OYW1lfX0ge3skdi5UeXBlTmFtZX19e3tpZiBndCAkc2l6ZSAkaX19LCB7e2VuZH19e3tlbmR9fSkgKnt7LlR5cGVOYW1lfX0gewogICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSByYW5nZSAuRmllbGRzfX0KICAgICAgICAgICAgICAgICAgICAgICAgICB7eyRJbnN0YW5jZVZhcn19Lnt7Lk5hbWV9fSA9IHt7Zmlyc3RMb3dlciAuTmFtZX19CiAgICAgICAgICAgICAgICAgICAgICAgICAge3stIGVuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICAgcmV0dXJuIHt7JEluc3RhbmNlVmFyfX0KICAgICAgICAgICAgICAgICAgICAgIH0sIGJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLk5hbWVzKFtdc3RyaW5newogICAgICAgICAgICAgICAgICAgICAgICAge3stIHJhbmdlIC5GaWVsZHN9fQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ7ey5BdXRvd2lyZWRNYXJrZXIuTmFtZX19e3tpZiBub3QgLkF1dG93aXJlZE1hcmtlci5SZXF1aXJlZH19LG9taXRlcnJvcnt7ZW5kfX0iLAogICAgICAgICAgICAgICAgICAgICAgICAge3stIGVuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICB9KSwKICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUG9zdENvbnN0cnVjdE1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUHJlQWZ0ZXJTZXQoInt7JFBvc3RDb25zdHJ1Y3RNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICAgICAgIHt7LSBpZiAkUHJlRGVzdHJveU1ldGhvZH19YmVhbi5DdXN0b21CZWFuRmFjdG9yeU9wdHMuUG9zdERlc3Ryb3koInt7JFByZURlc3Ryb3lNZXRob2R9fSIpLHt7ZW5kfX0KICAgICAgICAgICAgICAgICAgICAgICAgICkpKQogICAgICAgICAgICB7ey0gZWxzZX19CiAgICAgICAgICAgICAgICB7ey0gaWYgb3IgJFBvc3RDb25zdHJ1Y3RNZXRob2QgJFByZURlc3Ryb3lNZXRob2R9fQogICAgICAgICAgICAgICAgICAgIG5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oYmVhbi5OZXdDdXN0b21CZWFuRmFjdG9yeVdpdGhPcHRzKGZ1bmMoKSAqe3suVHlwZU5hbWV9fSB7IHJldHVybiB7eyRJbnN0YW5jZVZhcn19IH0sCiAgICAgICAgICAgICAgICAgICAgICAgICB7ey0gaWYgJFBvc3RDb25zdHJ1Y3RNZXRob2R9fWJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLlByZUFmdGVyU2V0KCJ7eyRQb3N0Q29uc3RydWN0TWV0aG9kfX0iKSx7e2VuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICB7ey0gaWYgJFByZURlc3Ryb3lNZXRob2R9fWJlYW4uQ3VzdG9tQmVhbkZhY3RvcnlPcHRzLlBvc3REZXN0cm95KCJ7eyRQcmVEZXN0cm95TWV0aG9kfX0iKSx7e2VuZH19CiAgICAgICAgICAgICAgICAgICAgICAgICApKSkKICAgICAgICAgICAgICAgIHt7LSBlbHNlfX0KICAgICAgICAgICAgICAgICAgICBuZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKHt7JEluc3RhbmNlVmFyfX0pKQogICAgICAgICAgICAgICAge3stIGVuZH19CiAgICAgICAgICAgIHt7LSBlbmR9fQogICAgICAgIHt7LSBlbmR9fQogICAge3stIGVuZH19Cnt7LSBlbmR9fQoKe3stIGlmIC5CZWFuTWFya2VyfX0Ke3stIGlmIC5CZWFuTWFya2VyLlZhbHVlfX0KCS8vIHJlZ2lzdGVyIHt7LlR5cGVOYW1lfX0KCXt7LSBpZiBwcm90b3R5cGUgLlNjb3BlTWFya2VyfX0KCXt7LSBpZiBvciAuQmVhbk1hcmtlci5Jbml0TWV0aG9kIC5CZWFuTWFya2VyLkRlc3Ryb3lNZXRob2R9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbkJ5TmFtZSgie3suQmVhbk1hcmtlci5WYWx1ZX19IiwgYmVhbi5OZXdDdXN0b21CZWFuRmFjdG9yeSh7ey5UeXBlTmFtZX19LCAie3suQmVhbk1hcmtlci5Jbml0TWV0aG9kfX0iLCAie3suQmVhbk1hcmtlci5EZXN0cm95TWV0aG9kfX0iKSkpCgl7ey0gZWxzZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5CZWFuTWFya2VyLlZhbHVlfX0iLCB7ey5UeXBlTmFtZX19KSkKCXt7LSBlbmR9fQoJe3stIGVsc2V9fQoJe3stIGlmIG9yIC5CZWFuTWFya2VyLkluaXRNZXRob2QgLkJlYW5NYXJrZXIuRGVzdHJveU1ldGhvZH19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5CZWFuTWFya2VyLlZhbHVlfX0iLCBiZWFuLk5ld0N1c3RvbUJlYW5GYWN0b3J5KGJlYW4uU2luZ2xldG9uRmFjdG9yeSh7ey5UeXBlTmFtZX19KSwgInt7LkJlYW5NYXJrZXIuSW5pdE1ldGhvZH19IiwgInt7LkJlYW5NYXJrZXIuRGVzdHJveU1ldGhvZH19IikpKQoJe3stIGVsc2V9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbkJ5TmFtZSgie3suQmVhbk1hcmtlci5WYWx1ZX19IiwgYmVhbi5TaW5nbGV0b25GYWN0b3J5KHt7LlR5cGVOYW1lfX0pKSkKCXt7LSBlbmR9fQoJe3stIGVuZH19Cnt7LSBlbHNlfX0KCS8vIHJlZ2lzdGVyIHt7LlR5cGVOYW1lfX0KCXt7LSBpZiBwcm90b3R5cGUgLlNjb3BlTWFya2VyfX0KCXt7LSBpZiBvciAuQmVhbk1hcmtlci5Jbml0TWV0aG9kIC5CZWFuTWFya2VyLkRlc3Ryb3lNZXRob2R9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbihiZWFuLk5ld0N1c3RvbUJlYW5GYWN0b3J5KHt7LlR5cGVOYW1lfX0sICJ7ey5CZWFuTWFya2VyLkluaXRNZXRob2R9fSIsICJ7ey5CZWFuTWFya2VyLkRlc3Ryb3lNZXRob2R9fSIpKSkKCXt7LSBlbHNlfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oe3suVHlwZU5hbWV9fSkpCgl7ey0gZW5kfX0KCXt7LSBlbHNlfX0KCXt7LSBpZiBvciAuQmVhbk1hcmtlci5Jbml0TWV0aG9kIC5CZWFuTWFya2VyLkRlc3Ryb3lNZXRob2R9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbihiZWFuLk5ld0N1c3RvbUJlYW5GYWN0b3J5KGJlYW4uU2luZ2xldG9uRmFjdG9yeSh7ey5UeXBlTmFtZX19KSwgInt7LkJlYW5NYXJrZXIuSW5pdE1ldGhvZH19IiwgInt7LkJlYW5NYXJrZXIuRGVzdHJveU1ldGhvZH19IikpKQoJe3stIGVsc2V9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbihiZWFuLlNpbmdsZXRvbkZhY3Rvcnkoe3suVHlwZU5hbWV9fSkpKQoJe3stIGVuZH19Cgl7ey0gZW5kfX0Ke3stIGVuZH19Cnt7LSBlbmR9fQoKe3stIHJhbmdlIC5NZXRob2RzIH19Cnt7aWYgLkJlYW5NYXJrZXIgfX0Ke3tpZiAuQmVhbk1hcmtlci5WYWx1ZX19CgkvLyByZWdpc3RlciB7eyRUeXBlTmFtZX19Lnt7Lk5hbWV9fQoJe3stIGlmIHByb3RvdHlwZSAuU2NvcGVNYXJrZXJ9fQoJe3stIGlmIG9yIC5CZWFuTWFya2VyLkluaXRNZXRob2QgLkJlYW5NYXJrZXIuRGVzdHJveU1ldGhvZH19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5CZWFuTWFya2VyLlZhbHVlfX0iLCBiZWFuLk5ld0N1c3RvbUJlYW5GYWN0b3J5KHt7JEluc3RhbmNlVmFyfX0ue3suTmFtZX19LCAie3suQmVhbk1hcmtlci5Jbml0TWV0aG9kfX0iLCAie3suQmVhbk1hcmtlci5EZXN0cm95TWV0aG9kfX0iKSkpCgl7ey0gZWxzZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5CZWFuTWFya2VyLlZhbHVlfX0iLCB7eyRJbnN0YW5jZVZhcn19Lnt7Lk5hbWV9fSkpCgl7ey0gZW5kfX0KCXt7LSBlbHNlfX0KCXt7LSBpZiBvciAuQmVhbk1hcmtlci5Jbml0TWV0aG9kIC5CZWFuTWFya2VyLkRlc3Ryb3lNZXRob2R9fQoJbmV2ZXJyb3IuUGFuaWNFcnJvcihib290LlJlZ2lzdGVyQmVhbkJ5TmFtZSgie3suQmVhbk1hcmtlci5WYWx1ZX19IiwgYmVhbi5OZXdDdXN0b21CZWFuRmFjdG9yeShiZWFuLlNpbmdsZXRvbkZhY3Rvcnkoe3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0pLCAie3suQmVhbk1hcmtlci5Jbml0TWV0aG9kfX0iLCAie3suQmVhbk1hcmtlci5EZXN0cm95TWV0aG9kfX0iKSkpCgl7ey0gZWxzZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuQnlOYW1lKCJ7ey5CZWFuTWFya2VyLlZhbHVlfX0iLCBiZWFuLlNpbmdsZXRvbkZhY3Rvcnkoe3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0pKSkKCXt7LSBlbmR9fQoJe3stIGVuZH19Cnt7ZWxzZX19CgkvLyByZWdpc3RlciB7eyRUeXBlTmFtZX19Lnt7Lk5hbWV9fQoJe3stIGlmIHByb3RvdHlwZSAuU2NvcGVNYXJrZXJ9fQoJe3stIGlmIG9yIC5CZWFuTWFya2VyLkluaXRNZXRob2QgLkJlYW5NYXJrZXIuRGVzdHJveU1ldGhvZH19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKGJlYW4uTmV3Q3VzdG9tQmVhbkZhY3Rvcnkoe3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0sICJ7ey5CZWFuTWFya2VyLkluaXRNZXRob2R9fSIsICJ7ey5CZWFuTWFya2VyLkRlc3Ryb3lNZXRob2R9fSIpKSkKCXt7LSBlbHNlfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oe3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0pKQoJe3stIGVuZH19Cgl7ey0gZWxzZX19Cgl7ey0gaWYgb3IgLkJlYW5NYXJrZXIuSW5pdE1ldGhvZCAuQmVhbk1hcmtlci5EZXN0cm95TWV0aG9kfX0KCW5ldmVycm9yLlBhbmljRXJyb3IoYm9vdC5SZWdpc3RlckJlYW4oYmVhbi5OZXdDdXN0b21CZWFuRmFjdG9yeShiZWFuLlNpbmdsZXRvbkZhY3Rvcnkoe3skSW5zdGFuY2VWYXJ9fS57ey5OYW1lfX0pLCAie3suQmVhbk1hcmtlci5Jbml0TWV0aG9kfX0iLCAie3suQmVhbk1hcmtlci5EZXN0cm95TWV0aG9kfX0iKSkpCgl7ey0gZWxzZX19CgluZXZlcnJvci5QYW5pY0Vycm9yKGJvb3QuUmVnaXN0ZXJCZWFuKGJlYW4uU2luZ2xldG9uRmFjdG9yeSh7eyRJbnN0YW5jZVZhcn19Lnt7Lk5hbWV9fSkpKQoJe3stIGVuZH19Cgl7ey0gZW5kfX0Ke3tlbmR9fQp7e2VuZH19Cnt7LSBlbmR9fQ==`
}

func getBuildTemplate(name string) string {
	d, err := base64.StdEncoding.DecodeString(buildinTemplate[name])
	if err != nil {
		return ""
	}
	return string(d)
}