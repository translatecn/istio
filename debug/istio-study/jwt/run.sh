#生成一个 JWK，通过模板指定 kid 为 youdianzhishi-key:
jwx jwk generate --keysize 4096 --type RSA --template '{"kid":"youdianzhishi-key"}' -o rsa.jwk

#从 rsa.jwk 中提取 JWK 公钥:
jwx jwk fmt --public-key -o rsa-public.jwk rsa.jwk

#将它们转换成 PEM 格式的公钥和私钥:
jwx jwk fmt -I json -O pem rsa.jwk
jwx jwk fmt -I json -O pem rsa-public.jwk

exp=$(perl -e 'print time() + 31536000, "\n"')
#签发 JWT Token
jwx jws sign --key rsa.jwk --alg RS256 --header '{"typ":"JWT"}' -o token.txt - <<EOF
{
  "iss": "testing@secure.istio.io",
  "sub": "cnych001",
  "iat": 1700648397,
  "exp": ${exp},
  "name": "Yang Ming"
}
EOF

python3 rep.py

#验证有效性
jwx jws verify --alg RS256 --key rsa-public.jwk token.txt
#iss: issuer，token 是谁签发的
#sub: token 的主体信息，一般设置为 token 代表用户身份的唯一 id 或唯一用户名

#{
#    "alg": "RS256",  # 算法「可选参数」
#    "kty": "RSA",    # 密钥类型
#    "use": "sig",    # 被用于签名「可选参数」
#    "kid": "DHFxxxxx_-envvQ",  # key 的唯一 id
#    "n": "xAExxxxMQ", 公钥的指数(exponent)
#    "e": "AQAB"  # 公钥的模数(modulus)
#}
