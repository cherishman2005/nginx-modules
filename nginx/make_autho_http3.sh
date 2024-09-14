./configure --prefix=/data/services/nginx/openresty --with-debug --with-http_v2_module --with-http_v3_module --with-stream \
            --with-http_stub_status_module \
            --with-cc-opt="-g -O2" \
            --with-openssl="../openssl-3.3.1"
 

make -j2 &&  make install

exit 0
