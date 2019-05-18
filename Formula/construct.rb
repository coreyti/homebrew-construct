require 'yaml'

class Construct < Formula
  VERSION = '0.1.1'

  desc     "Base functionality for Construct"
  homepage "https://github.com/coreyti/homebrew-construct/"
  url      "https://raw.githubusercontent.com/coreyti/homebrew-construct/master/Package/construct-#{VERSION}.tgz"
  sha256   "f35248d789b2ff27b7fec026dff4292e582d14245a9f21438c5607fc806d0dc0"

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
