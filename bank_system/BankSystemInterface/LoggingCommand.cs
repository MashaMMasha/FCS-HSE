namespace BankSystem;

public class LoggingCommand<T> : ICommand<T>
{
    private readonly ICommand<T> _command;
    private readonly string _description;

    public LoggingCommand(ICommand<T> command, string description)
    {
        _command = command;
        _description = description;
    }

    public T Execute()
    {
        T result = _command.Execute();
        if (result is bool)
        {
            bool boolResult = (bool)(object)result;
            if (boolResult == false)
            {
                Logger.Log($"Произошла ошибка при попытке выполнить операцию: {_description}", LogLevel.Error);
                return result;
            }
        }
        Logger.Log(_description,LogLevel.Info);
        return result;
    }
}
