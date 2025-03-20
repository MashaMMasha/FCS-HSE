using System;
using System.Collections.Generic;

namespace BankSystem;

public abstract class DataExporter : ExportNotifier
{
    public void Export(string filePath, List<BankAccount> accounts, List<Category> categories,
        List<Operation> operations)
    {
        try
        {
            NotifyExportStarted();

            string data = SerializeData(accounts, categories, operations);

            WriteFile(filePath, data);

            NotifyExportCompleted();
        }
        catch (Exception ex)
        {
            NotifyExportError($"Ошибка при экспорте: {ex.Message}");
        }
    }

    protected abstract string SerializeData(List<BankAccount> accounts, List<Category> categories,
        List<Operation> operations);

    private void WriteFile(string filePath, string data)
    {
        System.IO.File.WriteAllText(filePath, data);
        Console.WriteLine($"Данные экспортированы в файл: {filePath}");
    }
}