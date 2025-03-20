namespace BankSystem;

public class ConsoleObserver : IObserver
{
    public void OnExportStarted()
    {
        Console.WriteLine("Экспорт начат.");
    }

    public void OnExportCompleted()
    {
        Console.WriteLine("Экспорт завершен.");
    }

    public void OnExportError(string message)
    {
        Console.WriteLine($"Ошибка экспорта: {message}");
    }
}