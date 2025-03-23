namespace Preservary.Configuration.Models;
public class Location
{
    public string? SourcePath { get; set; }

    public string? DestinationPath { get; set; }

    public string? ComposePath { get; set; }

    public string BackupInterval { get; set; } = "1d";

    public bool Archive { get; set; }
}
