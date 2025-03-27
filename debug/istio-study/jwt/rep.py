rsa = ''
with open("rsa-public.jwk", 'r', encoding='utf8') as f:
    for line in f.readlines():
        rsa += ' ' * 16 + line

with open("jwt-example-tmp.yaml", 'r', encoding='utf8') as f:
    data = f.read()

data = data.replace('JWK', rsa)
print(data)

with open("jwt-example.yaml", 'w', encoding='utf8') as f:
    f.write(data)
