#===================##===================##===================##===================##===================##===================##===================##===================#
#====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

#===================##===================##===================##===================##===================##===================##===================##===================#
#====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======#
#===================##===================##===================##===================##===================##===================##===================##===================#


#===================##===================##===================##===================##===================##===================##===================##===================#
#====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======#
#===================##===================##===================##===================##===================##===================##===================##===================#



#===================##===================##===================##===================##===================##===================##===================##===================#
#====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

.PHONY: commit
## commit: add and commit files to develop branch
commit:
	@ git add . && git commit

.PHONY: status
## status: show git status
status:
	@ git status -u

.PHONY: push
## push: push application to main branch
push:
	@ git push -u origin main

.PHONY: push-d
## push-d: push application to develop branch
push-d:
	@ git push -u origin develop

