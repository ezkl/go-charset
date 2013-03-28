#!/usr/bin/env python

# CP949 specs;
# http://ko.wikipedia.org/wiki/%EC%BD%94%EB%93%9C_%ED%8E%98%EC%9D%B4%EC%A7%80_949
# http://msdn.microsoft.com/ko-kr/goglobal/cc305154.aspx

import re

if __name__ == '__main__':
    ptn = re.compile(r'0x([A-F0-9]*)\s*0x([A-F0-9]*)')

    pageBegin = 0
    lastCode = 0

    page_buf = ""

    # XXX: download CP949.TXT from;
    # http://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/WINDOWS/CP949.TXT
    for line in open("CP949.TXT"):
        if line.startswith("#"):
            continue
        rst = ptn.findall(line)
        if len(rst) != 1:
            continue

        curr_native, curr_unicode = int(rst[0][0], 16), int(rst[0][1], 16)
        if curr_native == curr_unicode:
            continue

        page_buf += unichr(curr_unicode).encode('utf-8')

        if curr_native - 1 != lastCode:
            if pageBegin == 0:
                pageBegin = curr_native
                page_buf += unichr(curr_unicode).encode('utf-8')
            else:
                print "0x%04x %d"%(pageBegin, lastCode-pageBegin+1)
                pageBegin = curr_native
                print page_buf
                page_buf = ""

        lastCode = curr_native

    print "0x%04x %d"%(pageBegin, lastCode-pageBegin+1)
    print page_buf
