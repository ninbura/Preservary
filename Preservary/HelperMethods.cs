using System.Text.RegularExpressions;

namespace Preservary;
public static partial class HelperMethods
{
    public static int ParseBackupInterval(string backupInterval)
    {
        var match = IntervalRegex().Match(backupInterval);

        if (!match.Success) throw new Exception($"Invalid backup interval: {backupInterval}");
        if (!int.TryParse(match.Groups[1].Value, out var value))
            throw new Exception($"Invalid backup interval value: {match.Groups[1].Value}");

        var unit = match.Groups[2].Value;

        var minutes = unit switch
        {
            "m" => value,
            "h" => value * 60,
            "d" => value * 60 * 24,
            _ => throw new Exception($"Invalid backup interval unit: {unit}")
        };

        return minutes;
    }

    [GeneratedRegex(@"^(\d+)([mhd])$")]
    private static partial Regex IntervalRegex();
}
