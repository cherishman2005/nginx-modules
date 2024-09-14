./configure --prefix=/data/services/nginx/openresty --with-debug --with-http_v2_module --with-stream \
            --with-http_stub_status_module \
            --with-cc-opt="-g -O2" \
            --with-openssl="../openssl-1.0.2k"
 

make -j2 &&  make install

exit 0

