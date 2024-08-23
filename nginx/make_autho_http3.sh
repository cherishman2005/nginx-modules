./configure --prefix=/data/services/nginx/openresty --with-debug --with-http_v2_module --with-http_v3_module --with-stream \
            --with-cc-opt="-g -O2" \
            --with-openssl="../openssl-1.0.2k"
 

make -j40 &&  make install

exit 0
