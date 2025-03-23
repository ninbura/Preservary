namespace Preservary;
public class LocationService
{
    public LocationService(IOptions<AppSettings> appSettings)
    {
        AppSettings = appSettings;
        EvaluateLocations();
    }

    private IOptions<AppSettings> AppSettings { get; }

    private void EvaluateLocations()
    {
        if (AppSettings.Value.Locations == null)
        {
            throw new Exception("No locations found in configuration.");
        }

        foreach (var location in AppSettings.Value.Locations)
        {
            _ = EvaluateLocation(location);
        }
    }

    private async Task EvaluateLocation(Location location)
    {

        var minutes = HelperMethods.ParseBackupInterval(location.BackupInterval);
        await Task.Delay(TimeSpan.FromMinutes(minutes));
        await EvaluateLocation(location);
    }
}
