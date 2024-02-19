docker image build -f Dockerfile -t forum .

docker run --detach -p 1234:4000 forum