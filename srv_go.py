import sys, time
import logging
import c_go as go

glob = 0
def handle(req):
	global glob
	glob = glob + 1
	#print "s: " + req.s
	if False:
		return None
	return go.SrvObj(5, "cinco")

if __name__ == '__main__':
	logging.basicConfig()
	go.serve(":8080", handle)
