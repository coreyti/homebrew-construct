require 'yaml'

class Construct < Formula
  VERSION = '0.1.0'

  desc     "Base functionality for Construct"
  homepage "https://github.com/coreyti/homebrew-construct/"
  url      "https://raw.githubusercontent.com/coreyti/homebrew-construct/master/Package/construct-#{VERSION}.tgz"
  sha256   "fa22ac7ca1936608553288172650ee10abc8afe833ed6c268ffa3edcc98c9451"

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
