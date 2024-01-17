FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# Install nodejs 20
USER root
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get install -y nodejs

USER vscode
# Install gow
RUN go install github.com/mitranim/gow@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

# Install lighthouse
USER root
RUN npm install -g lighthouse

# Install Chrome
RUN curl -sSL https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -o /tmp/chrome.deb
RUN apt-get install -y /tmp/chrome.deb
RUN rm /tmp/chrome.deb

USER vscode