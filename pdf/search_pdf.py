from glob import glob
import sys
import os
import re


def disjunct(*args):
	return '(?:%s)' % '|'.join(args)


def normalize(text):
    return re_space.sub(' ', text)


def to_str(b):
    return b.decode('utf-8', 'ignore')


def from_str(s):
    return s.encode('utf-8')


def get_text(path):
    with open(path, 'rb') as f:
        return f.read()


def get_contexts(text, n=10):
    contexts = []
    for m in re_search.finditer(text):
        ctx = text[m.start() - n:m.end() + n]
        contexts.append(normalize(to_str(ctx)))
    return contexts


term = sys.argv[1]
pattern = sys.argv[2:]
term = br'\b%s\b' % from_str(disjunct('williams', 'smith', 'jones'))
pattern = sys.argv[1:]


re_search = re.compile(term, re.MULTILINE | re.DOTALL | re.IGNORECASE)
re_space = re.compile(r'\s+', re.MULTILINE | re.DOTALL | re.IGNORECASE)

# print('pattern="%s"' % pattern)


files = []
for p in pattern:
    # print('$$', p)
    if os.path.isdir(p):
        for f in glob(p):
            # print('!!', f)
            files.append(f)
    else:
        files.append(p)
print('%d files' % len(files))

all_contexts = set()
for i, path in enumerate(files):
    text = get_text(path)
    contexts = get_contexts(text)
    if contexts:
        # print('%s %d bytes %d matches %s' % (path, len(text), len(contexts), contexts))
        all_contexts = all_contexts.union(contexts)

print('=' * 80)
for i, s in enumerate(sorted(all_contexts)):
    print('%3d: "%s"' % (i, s))
