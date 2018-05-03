# prepare to move to separate branch

.PHONY: pull-template
pull-template:
	rm -rf template
	git clone git@bitbucket.org:octotpl/materialkit.git template
	( \
	  cd template && \
	  git checkout 2.x && \
	  git describe --tags \
	)
	( \
	  cp -n template/Template/assets/css/material-kit.min.css docs/assets/css/ && \
	  cp -n template/Template/assets/js/bootstrap-material-design.min.js docs/assets/js/ \
	  cp -n template/Template/assets/js/material-kit.min.js docs/assets/js/ \
	)
	rm -rf template/.git
