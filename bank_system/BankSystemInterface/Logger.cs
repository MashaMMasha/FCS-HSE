namespace BankSystem;

enum LogLevel
{
    Error,
    Warning,
    Info
}
public class Logger
{
    public static void Log(string message, LogLevel level)
    {
        switch (level)
        {
            case LogLevel.Error:
                Console.WriteLine($"[ERROR] {DateTime.Now}: {message}");
                break;
            case LogLevel.Warning:
                Console.WriteLine($"[WARNING] {DateTime.Now}: {message}");
                break;
            case LogLevel.Info:
                Console.WriteLine($"[INFO] {DateTime.Now}: {message}");
                break;
            default:
                Console.WriteLine($"[LOG] {DateTime.Now}: {message}");
                break;
        }
    }
    
}