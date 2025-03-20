namespace BankSystem;
public class BankCommand<T> : ICommand<T>
{
    private readonly Func<T> _func;

    public BankCommand(Func<T> func)
    {
        _func = func;
    }

    public T Execute()
    {
        return _func.Invoke();
    }
}
