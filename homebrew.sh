#!/bin/bash -xe

dist_bin="$1"
dist_sha256=$(shasum -a 256 -b "$dist_bin" | cut -f1 -d' ')
cat <<EOF
class Ghcp < Formula
  desc "Copy files to a repository using GitHub API"
  homepage "https://github.com/int128/ghcp"
  url "https://github.com/int128/ghcp/releases/download/${CIRCLE_TAG}/ghcp_darwin_amd64"
  version "${CIRCLE_TAG}"
  sha256 "${dist_sha256}"

  def install
    bin.install "ghcp_darwin_amd64" => "ghcp"
  end

  test do
    system "#{bin}/ghcp -h"
  end
end
EOF
