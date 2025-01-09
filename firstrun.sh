if [ $# -ne 2 ]; then
  echo "Usage: $0 <user> <pwd>"
  exit 1
fi

user=$1
pwd=$2

mkdir -p ./easy/db/

exit 1

docker run \
-e "DATABASE_LOC=/app/db/blog.db" \
-p 8080:8080 \
-v ./easy/db/:/app/db/:rw easyblog-app \
--user=$user --pwd=$pwd
