namespace BankSystem;

public interface ICommand<out T>
{
    T Execute();
}