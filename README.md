# project-p-back

+ 1 -> baixar golang versão 1.20
+ 2 -> baixar repo
+ 3 -> abrir terminal no diretório base
+ 4 -> rodar comando "go run main.go"

## endpoints
todos ficam dentro do api/route.go

# api/user -> post para criar nova conta
exemplo do body:
```json
{
	"username" : "rodrigo_test9",
	"password" : "12345"
}
```

exemplo do response correto:
```json
{
	"code": 201,
	"data": {
		"username": "rodrigo_test9",
		"password": "$2a$10$lGhmtbwOGOy/YmkY3aZoXe/PzvGwz/wzYKFzHKfnAA11uy5/j09Fy"
	},
	"message": "Created"
}
```
# api/user/login ->post do login para validar acesso e conceder token
exemplo do body:
```json
{
	"username" : "rodrigo_test1",
	"password" : "12345"
}
```

exemplo do response correto:
```json
{
	"code": 200,
	"data": {
		"acess_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTE0MjE0MjUsInVzZXJfaWQiOiI2NGNlZTBjNGQyMzNlYjkyZTYwN2RkODgifQ.d51rSXfrVHjqNqpdlcXQUSIIHYclpAadq-KRhwdIj8U",
		"expired": 1691421425,
		"user_id": "64cee0c4d233eb92e607dd88"
	},
	"message": "OK"
}
```

# api/graphql -> post query para consumir chatGpt 3.5 ( nescessário fornecer bearer token recebido do login)
exemplo do body:
```graphql
query {
	GPT3dot5(message: "Exemplo de mensagem generica", user_id: "64cee0c4d233eb92e607dd88") {
		id
		object
		created
		model
		usage {
			prompt_tokens
			completion_tokens
			total_tokens
		}
		choices {
			index
			message {
				role
				content
			}
			finish_reason
		}
	}
}
```

exemplo de response:
```json
{
	"code": 200,
	"data": {
		"data": {
			"GPT3dot5": {
				"choices": [
					{
						"finish_reason": "stop",
						"index": 0,
						"message": {
							"content": "Caro(a) Senhor(a),\n\nEspero que esta mensagem o encontre bem. Eu gostaria de lhe agradecer por seu interesse em nossa empresa/produto/serviço. Estamos encantados em saber que você está interessado em saber mais sobre o que temos a oferecer.\n\nNossa empresa/produto/serviço é reconhecido por [inserir características relevantes, como qualidade, inovação, confiabilidade], o que nos torna uma escolha confiável para atender às suas necessidades. Acreditamos que podemos ultrapassar suas expectativas e fornecer uma solução sob medida para você.\n\nGostaríamos de convidá-lo a [opção 1: visitar nosso site/oficina/showroom], onde você poderá explorar nosso portfólio, obter mais informações e entrar em contato com nossa equipe para esclarecer quaisquer dúvidas que você possa ter. Teremos prazer em ajudá-lo e garantir que você tenha uma excelente experiência com nossa empresa.\n\nSempre nos esforçamos para fornecer um atendimento de primeira linha e soluções personalizadas para cada cliente. Valorizamos sua opinião e gostaríamos de saber mais sobre suas necessidades específicas. Por favor, não hesite em entrar em contato conosco para agendar uma reunião, receber uma cotação ou qualquer outra informação que você precise.\n\nAgradecemos novamente por considerar nossa empresa/produto/serviço e esperamos ter a oportunidade de trabalhar com você em breve.\n\nAtenciosamente,\n\n[Seu nome]\n[Seu cargo]\n[Nome da empresa]",
							"role": "assistant"
						}
					}
				],
				"created": 1691420717,
				"id": "chatcmpl-7kw5FV6hgV9wBptQmenkq9Z4Lu4EU",
				"model": "gpt-3.5-turbo-0613",
				"object": "chat.completion",
				"usage": {
					"completion_tokens": 363,
					"prompt_tokens": 13,
					"total_tokens": 376
				}
			}
		}
	},
	"message": "OK"
}
```

