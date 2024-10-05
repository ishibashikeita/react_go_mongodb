# ベースイメージとしてnodeを使う 
FROM node:16-alpine  
# 作業ディレクトリを設定  
WORKDIR /usr/src/app    
# package.jsonとyarn.lockをコピーする  
COPY package*.json ./   
COPY yarn.lock ./     
# 依存しているパッケージのインストールをする   
RUN yarn install     
# 全ファイルをコピー   
COPY . .     
# 起動   
CMD [ "yarn", "start" ]