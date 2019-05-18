require 'yaml'

class Construct < Formula
  VERSION = '0.1.2'

  desc     "Base functionality for Construct"
  homepage "https://github.com/coreyti/homebrew-construct/"
  url      "https://raw.githubusercontent.com/coreyti/homebrew-construct/master/Package/construct-#{VERSION}.tgz"
  sha256   "8cbf3cc087aa4a118f19866864ec724e2e9029fb4561d258ebaaf5d8f8f8c440"

  def install
    bin.install "bin/construct"

    config["version"] = VERSION
    pkgshare.join("bin").mkpath

    pkgshare.join("config").write(config.to_yaml)
  end

  private

  def config
    @config ||= begin
      path = share.join("config")
      path.exist? ? YAML.load(File.read(path)) : {}
    end
  end
end
