# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Imgcat < Formula
  desc "Golang implementation of the Inline Images Protocol"
  homepage ""
  url "https://github.com/woodfairy/imgcat/releases/download/v0.1/imgcat.zip"
  sha256 "1b08d74945a80924e74a891e3a6345704d6c1b8061a6c0dafd341c333217a4a0"
  license ""

  # depends_on "cmake" => :build

  def install
  bin.install "imgcat"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test imgcat`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
