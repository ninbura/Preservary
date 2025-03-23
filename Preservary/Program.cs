using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Preservary.Configuration.Models;

var host = Host.CreateDefaultBuilder()
    .ConfigureHostConfiguration(hostOptions =>
    {
        hostOptions.SetBasePath(Directory.GetCurrentDirectory());
    })
    .ConfigureAppConfiguration((hostContext, appConfiguration) =>
    {
        appConfiguration.AddJsonFile("appsettings.json", optional: false);
    })
    .ConfigureServices((hostContext, services) =>
    {
        services.Configure<AppSettings>(hostContext.Configuration);
        services.AddSingleton<App>();
    })
    .Build();

await host.Services.GetRequiredService<App>().Run();
await host.RunAsync();