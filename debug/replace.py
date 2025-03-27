import os

rs = [
    ['kiwigrid/k8s-sidecar', 'registry.cn-hangzhou.aliyuncs.com/acejilam/k8s-sidecar'],
    ['ghcr.io/prometheus-operator/', 'registry.cn-hangzhou.aliyuncs.com/acejilam/'],
    ['prom/prometheus', 'registry.cn-hangzhou.aliyuncs.com/acejilam/prometheus'],
    ['hub: gcr.io/istio-testing', 'hub: registry.cn-hangzhou.aliyuncs.com/acejilam'],
    ['docker.io/istio/', 'registry.cn-hangzhou.aliyuncs.com/acejilam/'],
    ['docker.io/istio', 'registry.cn-hangzhou.aliyuncs.com/acejilam'],
    ['image: apache/', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/'],
    ['docker.io/jaegertracing', 'registry.cn-hangzhou.aliyuncs.com/acejilam'],
    ['docker.io/grafana', 'registry.cn-hangzhou.aliyuncs.com/acejilam'],
    ['docker.io/mccutchen', 'registry.cn-hangzhou.aliyuncs.com/acejilam'],
    ['image: hiroakis/', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/'],
    ['image: hiroakis/', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/'],
    ['image: curlimages/curl', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/curl'],
    ['image: busybox', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/busybox'],
    ['image: ghcr.io/spiffe/', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/'],
    ['image: registry.k8s.io/sig-storage/', 'image: registry.cn-hangzhou.aliyuncs.com/acejilam/'],
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
                if 'install.istio.io/v1alpha1' in data and 'hub: registry.cn-hangzhou.aliyuncs.com/acejilam' not in data:
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
                            f.write('  hub: registry.cn-hangzhou.aliyuncs.com/acejilam\n')
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
