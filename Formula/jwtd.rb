class Jwtd < Formula
  desc "A simple command-line JSON Web Tokens decoder tool"
  homepage "https://github.com/dandehoon/jwtd"
  license "MIT"

  if OS.mac?
    if Hardware::CPU.arm?
      url "https://github.com/dandehoon/jwtd/releases/latest/download/jwtd-darwin-arm64", using: :nounzip
    else
      url "https://github.com/dandehoon/jwtd/releases/latest/download/jwtd-darwin-amd64", using: :nounzip
    end
  elsif OS.linux?
    if Hardware::CPU.arm?
      url "https://github.com/dandehoon/jwtd/releases/latest/download/jwtd-linux-arm64", using: :nounzip
    else
      url "https://github.com/dandehoon/jwtd/releases/latest/download/jwtd-linux-amd64", using: :nounzip
    end
  end

  # Skip SHA256 validation for latest releases
  sha256 :no_check

  def install
    bin.install Dir["jwtd*"].first => "jwtd"
  end

  def caveats
    <<~EOS
      This formula installs the latest release automatically.
      Run 'brew upgrade jwtd' to get newer versions.
    EOS
  end

  test do
    # Test that the binary was installed and is executable
    assert_predicate bin/"jwtd", :exist?
    assert_predicate bin/"jwtd", :executable?

    # Test basic functionality
    sample_jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    output = shell_output("#{bin}/jwtd #{sample_jwt}")
    assert_match "John Doe", output
  end
end
