#coding=utf-8
import os,hashlib
from multiprocessing import Pool
 
def get_pic_list():
    file_list = os.listdir('/Users/JingjingHe/go/src/girls/images')
    print len(file_list)
    for i in file_list:
        file_list[file_list.index(i)] = '/Users/JingjingHe/go/src/girls/images/'+i
    pool = Pool()
    result = pool.map(cal_md5,file_list)
    reset_dict = {}
    for i in result:
        if i[0] not in reset_dict:
            reset_dict[i[0]] = i[1]
    del_list = []
    for i in reset_dict:
        del_list.append(reset_dict[i])
    print len(del_list)
    for i in file_list:
        if i not in del_list:
            os.system('rm '+i)

def cal_md5(file_name):
    try:
        m = hashlib.md5()
        m.update(open(file_name).read())
        psw = m.hexdigest()
        return (psw,file_name)
    except:
        pass

get_pic_list()
        
    