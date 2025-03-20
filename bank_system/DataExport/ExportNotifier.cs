using System;
using System.Collections.Generic;

namespace BankSystem;

public class ExportNotifier
{
    private readonly List<IObserver> _observers = new List<IObserver>();

    public void AddObserver(IObserver observer)
    {
        _observers.Add(observer);
    }

    public void RemoveObserver(IObserver observer)
    {
        _observers.Remove(observer);
    }

    protected void NotifyExportStarted()
    {
        foreach (var observer in _observers)
        {
            observer.OnExportStarted();
        }
    }

    protected void NotifyExportCompleted()
    {
        foreach (var observer in _observers)
        {
            observer.OnExportCompleted();
        }
    }

    protected void NotifyExportError(string message)
    {
        foreach (var observer in _observers)
        {
            observer.OnExportError(message);
        }
    }
}