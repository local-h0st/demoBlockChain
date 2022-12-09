package dem0chain

func PackageInfo() string {
	return `basic package containing block struct, chain struct, pow, etc.
	貌似package名字和文件名一样的话会自动导入,
	但是导入完了改成不一样也不会有错.
	为了方便自动导入,
	第一个go文件名就和package一样好了,
	其他的文件名随意, 
	同一个目录下package名一样就行,
	另外最好是目录名等于package名,
	也是为了方便自动导入`

}
