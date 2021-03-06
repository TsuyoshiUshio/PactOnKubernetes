# -*- encoding: utf-8 -*-
# stub: pact-provider-proxy 2.1.0 ruby lib vendor/rack-reverse-proxy/lib

Gem::Specification.new do |s|
  s.name = "pact-provider-proxy".freeze
  s.version = "2.1.0"

  s.required_rubygems_version = Gem::Requirement.new(">= 0".freeze) if s.respond_to? :required_rubygems_version=
  s.require_paths = ["lib".freeze, "vendor/rack-reverse-proxy/lib".freeze]
  s.authors = ["Beth".freeze]
  s.date = "2015-04-21"
  s.description = "See summary".freeze
  s.email = ["beth@bethesque.com".freeze]
  s.homepage = "https://github.com/bethesque/pact-provider-proxy".freeze
  s.licenses = ["MIT".freeze]
  s.rubygems_version = "2.6.12".freeze
  s.summary = "Allows verification of a pact against a running provider".freeze

  s.installed_by_version = "2.6.12" if s.respond_to? :installed_by_version

  if s.respond_to? :specification_version then
    s.specification_version = 4

    if Gem::Version.new(Gem::VERSION) >= Gem::Version.new('1.2.0') then
      s.add_runtime_dependency(%q<pact>.freeze, [">= 1.1.1", "~> 1.0"])
      s.add_development_dependency(%q<bundler>.freeze, ["~> 1.3"])
      s.add_development_dependency(%q<rake>.freeze, [">= 0"])
      s.add_development_dependency(%q<pry>.freeze, [">= 0"])
      s.add_runtime_dependency(%q<rack>.freeze, [">= 1.0.0"])
      s.add_runtime_dependency(%q<rack-proxy>.freeze, ["~> 0.5"])
    else
      s.add_dependency(%q<pact>.freeze, [">= 1.1.1", "~> 1.0"])
      s.add_dependency(%q<bundler>.freeze, ["~> 1.3"])
      s.add_dependency(%q<rake>.freeze, [">= 0"])
      s.add_dependency(%q<pry>.freeze, [">= 0"])
      s.add_dependency(%q<rack>.freeze, [">= 1.0.0"])
      s.add_dependency(%q<rack-proxy>.freeze, ["~> 0.5"])
    end
  else
    s.add_dependency(%q<pact>.freeze, [">= 1.1.1", "~> 1.0"])
    s.add_dependency(%q<bundler>.freeze, ["~> 1.3"])
    s.add_dependency(%q<rake>.freeze, [">= 0"])
    s.add_dependency(%q<pry>.freeze, [">= 0"])
    s.add_dependency(%q<rack>.freeze, [">= 1.0.0"])
    s.add_dependency(%q<rack-proxy>.freeze, ["~> 0.5"])
  end
end
