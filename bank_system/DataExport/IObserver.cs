namespace BankSystem;

public interface IObserver
{
    void OnExportStarted();
    void OnExportCompleted();
    void OnExportError(string message);
}