# server.py
import uuid

from flask import Flask, request, jsonify

app = Flask(__name__)

endpoints_db = {}


@app.route('/endpoints', methods=['POST'])
def add_endpoint():
    data = request.get_json()
    endpoint_id = str(uuid.uuid4())
    endpoints_db[endpoint_id] = data
    return jsonify({"id": endpoint_id}), 201


# curl --location --request POST 'http://localhost:8080/endpoints' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "address": "172.17.0.7",
#     "port_value": 8081
# }'
# curl --location --request POST 'http://localhost:8080/endpoints' \
# --header 'Content-Type: application/json' \
# --data '{
#     "address": "172.17.0.7",
#     "port_value": 8081
# }'
# curl --location --request POST 'http://localhost:8080/endpoints' \
# > --header 'Content-Type: application/json' \
# > --data '{
# >     "address": "172.17.0.7",
# >     "port_value": 8081
# > }'

@app.route('/v3/discovery:endpoints', methods=['POST'])
def discovery_endpoints():
    xds_response = {  # 构造 xDS v3 EDS 响应格式
        "version_info": "0",
        "resources": [
            {
                "@type": "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment",
                "cluster_name": "localservices",
                "endpoints": [
                    {
                        "lb_endpoints": [
                            {
                                "endpoint": {
                                    "address": {
                                        "socket_address": endpoint
                                    }
                                }
                            }
                            for endpoint in endpoints_db.values()
                        ]
                    }
                ]
            }
        ],
        "type_url": "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment",
        "nonce": "0"
    }
    return jsonify(xds_response)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8080, debug=True)
