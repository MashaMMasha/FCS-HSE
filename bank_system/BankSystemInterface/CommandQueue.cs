namespace BankSystem;

public class CommandQueue
{
    public Queue<ICommand> commands = new Queue<ICommand>();
    public void AddCommand(ICommand command)
    {
        commands.Enqueue(command);
    }
    public void ProcessCommands()
    {
        while (commands.Count > 0)
        {
            ICommand command = commands.Dequeue();
            command.Execute();
        }
    }
}