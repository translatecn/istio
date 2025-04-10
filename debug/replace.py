import os

rs = [
    ['kiwigrid/k8s-sidecar', 'ccr.ccs.tencentyun.com/acejilam/k8s-sidecar'],
    ['ghcr.io/prometheus-operator/', 'ccr.ccs.tencentyun.com/acejilam/'],
    ['prom/prometheus', 'ccr.ccs.tencentyun.com/acejilam/prometheus'],
    ['hub: gcr.io/istio-testing', 'hub: ccr.ccs.tencentyun.com/acejilam'],
    ['docker.io/istio/', 'ccr.ccs.tencentyun.com/acejilam/'],
    ['docker.io/istio', 'ccr.ccs.tencentyun.com/acejilam'],
    ['image: apache/', 'image: ccr.ccs.tencentyun.com/acejilam/'],
    ['docker.io/jaegertracing', 'ccr.ccs.tencentyun.com/acejilam'],
    ['docker.io/grafana', 'ccr.ccs.tencentyun.com/acejilam'],
    ['docker.io/mccutchen', 'ccr.ccs.tencentyun.com/acejilam'],
    ['image: hiroakis/', 'image: ccr.ccs.tencentyun.com/acejilam/'],
    ['image: hiroakis/', 'image: ccr.ccs.tencentyun.com/acejilam/'],
    ['image: curlimages/curl', 'image: ccr.ccs.tencentyun.com/acejilam/curl'],
    ['image: busybox', 'image: ccr.ccs.tencentyun.com/acejilam/busybox'],
    ['image: ghcr.io/spiffe/', 'image: ccr.ccs.tencentyun.com/acejilam/'],
    ['image: registry.k8s.io/sig-storage/', 'image: ccr.ccs.tencentyun.com/acejilam/'],
    ['''  volumeClaimTemplates:
    - apiVersion: v1
      kind: PersistentVolumeClaim
      metadata:
        name: storage
      spec:
        accessModes:
          - ReadWriteOnce
        resources:
          requests:
            storage: "10Gi"''',
     '''        - name: storage
          emptyDir: {}
          '''],
]

ISTIO_PATH = '/Users/acejilam/Desktop/ebpf/istio'

for cd, _dirs, files in os.walk(ISTIO_PATH):
    for file in files:
        path = os.path.join(cd, file)
        if path.endswith('.sh') or path.endswith('.yml') or path.endswith('.yaml') or path.endswith('Dockerfile'):
            print(path)
            with open(path, 'r', encoding='utf8') as f:
                data = f.read()
                for item in rs:
                    data = data.replace(item[0], item[1])
            with open(path, 'w', encoding='utf8') as f:
                f.write(data)
for cd, _dirs, files in os.walk(ISTIO_PATH):
    for file in files:
        path = os.path.join(cd, file)
        if path.endswith('.yaml'):
            skip = True
            with open(path, 'r', encoding='utf8') as f:
                data = f.read()
                if 'install.istio.io/v1alpha1' in data and 'hub: ccr.ccs.tencentyun.com/acejilam' not in data:
                    skip = False
            if not skip:
                with open(path, 'w', encoding='utf8') as f:
                    install = False
                    spec = False
                    for line in data.split('\n'):
                        f.write(line + '\n')
                        if 'apiVersion: install.istio.io' in line:
                            install = True
                        if 'spec:' in line and install:
                            spec = True
                        if spec and install:
                            f.write('  hub: ccr.ccs.tencentyun.com/acejilam\n')
                            install = False
                            spec = False

with open(
        os.path.join(f"{ISTIO_PATH}", "samples/bookinfo/demo-profile-no-gateways.yaml"),
        "a", encoding='utf8') as f:
    f.write('\n')
    f.write('''  meshConfig:
    defaultProviders:
      tracing:
      - "skywalking"
    enableTracing: true
    extensionProviders:
    - name: "skywalking"
      skywalking:
        service: tracing.istio-system.svc.cluster.local
        port: 11800''')
