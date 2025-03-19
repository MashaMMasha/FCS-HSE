namespace BankSystem;

public class BankCommand : ICommand
{
    private readonly Action _action;
    private readonly string _description;
    public BankCommand(Action action, string description)
    {
        _action = action;
        _description = description;
    }

    public void Execute()
    {
        Logger.Log(_description, LogLevel.Info);
    }
}