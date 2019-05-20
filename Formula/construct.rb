require 'yaml'

class Construct < Formula
  VERSION = '0.1.3'

  desc     "Base functionality for Construct"
  homepage "https://github.com/coreyti/homebrew-construct/"
  url      "https://raw.githubusercontent.com/coreyti/homebrew-construct/master/Package/construct-#{VERSION}.tgz"
  sha256   "bd20aab4ec557cdfd3f3df3d08ef0f726ed3e8abfc16cf8003554193a592aacf"

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
