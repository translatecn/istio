import os

ns = set()

for cd, dirs, files in os.walk("./logs"):
    files.sort()
    for file in files:
        p = os.path.join('./logs', file)
        with open(p, 'r', encoding='utf8') as f:
            data = f.read()
        if data in ns:
            os.remove(p)
        else:
            ns.add(data)
